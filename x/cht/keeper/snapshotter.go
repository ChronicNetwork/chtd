package keeper

import (
	"encoding/hex"
	"io"

	snapshot "github.com/cosmos/cosmos-sdk/snapshots/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	protoio "github.com/gogo/protobuf/io"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/ChronicNetwork/cht/x/cht/ioutils"
	"github.com/ChronicNetwork/cht/x/cht/types"
)

var _ snapshot.ExtensionSnapshotter = &ChtSnapshotter{}

// SnapshotFormat format 1 is just gzipped cht byte code for each item payload. No protobuf envelope, no metadata.
const SnapshotFormat = 1

type ChtSnapshotter struct {
	cht *Keeper
	cms sdk.MultiStore
}

func NewChtSnapshotter(cms sdk.MultiStore, cht *Keeper) *ChtSnapshotter {
	return &ChtSnapshotter{
		cht: cht,
		cms: cms,
	}
}

func (ws *ChtSnapshotter) SnapshotName() string {
	return types.ModuleName
}

func (ws *ChtSnapshotter) SnapshotFormat() uint32 {
	return SnapshotFormat
}

func (ws *ChtSnapshotter) SupportedFormats() []uint32 {
	// If we support older formats, add them here and handle them in Restore
	return []uint32{SnapshotFormat}
}

func (ws *ChtSnapshotter) Snapshot(height uint64, protoWriter protoio.Writer) error {
	cacheMS, err := ws.cms.CacheMultiStoreWithVersion(int64(height))
	if err != nil {
		return err
	}

	ctx := sdk.NewContext(cacheMS, tmproto.Header{}, false, log.NewNopLogger())
	seenBefore := make(map[string]bool)
	var rerr error

	ws.cht.IterateCodeInfos(ctx, func(id uint64, info types.CodeInfo) bool {
		// Many code ids may point to the same code hash... only sync it once
		hexHash := hex.EncodeToString(info.CodeHash)
		// if seenBefore, just skip this one and move to the next
		if seenBefore[hexHash] {
			return false
		}
		seenBefore[hexHash] = true

		// load code and abort on error
		chtBytes, err := ws.cht.GetByteCode(ctx, id)
		if err != nil {
			rerr = err
			return true
		}

		compressedCht, err := ioutils.GzipIt(chtBytes)
		if err != nil {
			rerr = err
			return true
		}

		err = snapshot.WriteExtensionItem(protoWriter, compressedCht)
		if err != nil {
			rerr = err
			return true
		}

		return false
	})

	return rerr
}

func (ws *ChtSnapshotter) Restore(
	height uint64, format uint32, protoReader protoio.Reader,
) (snapshot.SnapshotItem, error) {
	if format == SnapshotFormat {
		return ws.processAllItems(height, protoReader, restoreV1, finalizeV1)
	}
	return snapshot.SnapshotItem{}, snapshot.ErrUnknownFormat
}

func restoreV1(ctx sdk.Context, k *Keeper, compressedCode []byte) error {
	chtCode, err := ioutils.Uncompress(compressedCode, uint64(types.MaxChtSize))
	if err != nil {
		return sdkerrors.Wrap(types.ErrCreateFailed, err.Error())
	}

	// FIXME: check which codeIDs the checksum matches??
	_, err = k.wasmVM.Create(chtCode)
	if err != nil {
		return sdkerrors.Wrap(types.ErrCreateFailed, err.Error())
	}
	return nil
}

func finalizeV1(ctx sdk.Context, k *Keeper) error {
	// FIXME: ensure all codes have been uploaded?
	return k.InitializePinnedCodes(ctx)
}

func (ws *ChtSnapshotter) processAllItems(
	height uint64,
	protoReader protoio.Reader,
	cb func(sdk.Context, *Keeper, []byte) error,
	finalize func(sdk.Context, *Keeper) error,
) (snapshot.SnapshotItem, error) {
	ctx := sdk.NewContext(ws.cms, tmproto.Header{Height: int64(height)}, false, log.NewNopLogger())

	// keep the last item here... if we break, it will either be empty (if we hit io.EOF)
	// or contain the last item (if we hit payload == nil)
	var item snapshot.SnapshotItem
	for {
		item = snapshot.SnapshotItem{}
		err := protoReader.ReadMsg(&item)
		if err == io.EOF {
			break
		} else if err != nil {
			return snapshot.SnapshotItem{}, sdkerrors.Wrap(err, "invalid protobuf message")
		}

		// if it is not another ExtensionPayload message, then it is not for us.
		// we should return it an let the manager handle this one
		payload := item.GetExtensionPayload()
		if payload == nil {
			break
		}

		if err := cb(ctx, ws.cht, payload.Payload); err != nil {
			return snapshot.SnapshotItem{}, sdkerrors.Wrap(err, "processing snapshot item")
		}
	}

	return item, finalize(ctx, ws.cht)
}
