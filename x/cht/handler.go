package cht

import (
	"fmt"

	"github.com/gogo/protobuf/proto"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/ChronicNetwork/chtd/x/cht/keeper"
	"github.com/ChronicNetwork/chtd/x/cht/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler returns a handler for "cht" type messages.
func NewHandler(k types.ContractOpsKeeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		var (
			res proto.Message
			err error
		)
		switch msg := msg.(type) {
		case *MsgStoreCode: //nolint:typecheck
			res, err = msgServer.StoreCode(sdk.WrapSDKContext(ctx), msg)
		case *MsgInstantiateContract:
			res, err = msgServer.InstantiateContract(sdk.WrapSDKContext(ctx), msg)
		case *MsgExecuteContract:
			res, err = msgServer.ExecuteContract(sdk.WrapSDKContext(ctx), msg)
		case *MsgMigrateContract:
			res, err = msgServer.MigrateContract(sdk.WrapSDKContext(ctx), msg)
		case *MsgUpdateAdmin:
			res, err = msgServer.UpdateAdmin(sdk.WrapSDKContext(ctx), msg)
		case *MsgClearAdmin:
			res, err = msgServer.ClearAdmin(sdk.WrapSDKContext(ctx), msg)
		default:
			errMsg := fmt.Sprintf("unrecognized cht message type: %T", msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}

		ctx = ctx.WithEventManager(filterMessageEvents(ctx))
		return sdk.WrapServiceResult(ctx, res, err)
	}
}

// filterMessageEvents returns the same events with all of type == EventTypeMessage removed except
// for cht message types.
// this is so only our top-level message event comes through
func filterMessageEvents(ctx sdk.Context) *sdk.EventManager {
	m := sdk.NewEventManager()
	for _, e := range ctx.EventManager().Events() {
		if e.Type == sdk.EventTypeMessage &&
			!hasChtModuleAttribute(e.Attributes) {
			continue
		}
		m.EmitEvent(e)
	}
	return m
}

func hasChtModuleAttribute(attrs []abci.EventAttribute) bool {
	for _, a := range attrs {
		if sdk.AttributeKeyModule == string(a.Key) &&
			types.ModuleName == string(a.Value) {
			return true
		}
	}
	return false
}
