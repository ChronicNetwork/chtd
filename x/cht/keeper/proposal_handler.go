package keeper

import (
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/ChronicToken/cht/x/cht/types"
)

// NewChtProposalHandler creates a new governance Handler for cht proposals
func NewChtProposalHandler(k decoratedKeeper, enabledProposalTypes []types.ProposalType) govtypes.Handler {
	return NewChtProposalHandlerX(NewGovPermissionKeeper(k), enabledProposalTypes)
}

// NewChtProposalHandlerX creates a new governance Handler for cht proposals
func NewChtProposalHandlerX(k types.ContractOpsKeeper, enabledProposalTypes []types.ProposalType) govtypes.Handler {
	enabledTypes := make(map[string]struct{}, len(enabledProposalTypes))
	for i := range enabledProposalTypes {
		enabledTypes[string(enabledProposalTypes[i])] = struct{}{}
	}
	return func(ctx sdk.Context, content govtypes.Content) error {
		if content == nil {
			return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "content must not be empty")
		}
		if _, ok := enabledTypes[content.ProposalType()]; !ok {
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unsupported chronic proposal content type: %q", content.ProposalType())
		}
		switch c := content.(type) {
		case *types.StoreCodeProposal:
			return handleStoreCodeProposal(ctx, k, *c)
		case *types.InstantiateContractProposal:
			return handleInstantiateProposal(ctx, k, *c)
		case *types.MigrateContractProposal:
			return handleMigrateProposal(ctx, k, *c)
		case *types.UpdateAdminProposal:
			return handleUpdateAdminProposal(ctx, k, *c)
		case *types.ClearAdminProposal:
			return handleClearAdminProposal(ctx, k, *c)
		case *types.PinCodesProposal:
			return handlePinCodesProposal(ctx, k, *c)
		case *types.UnpinCodesProposal:
			return handleUnpinCodesProposal(ctx, k, *c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized chronic proposal content type: %T", c)
		}
	}
}

func handleStoreCodeProposal(ctx sdk.Context, k types.ContractOpsKeeper, p types.StoreCodeProposal) error {
	if err := p.ValidateBasic(); err != nil {
		return err
	}

	runAsAddr, err := sdk.AccAddressFromBech32(p.RunAs)
	if err != nil {
		return sdkerrors.Wrap(err, "run as address")
	}
	_, err = k.Create(ctx, runAsAddr, p.WASMByteCode, p.InstantiatePermission)
	return err
}

func handleInstantiateProposal(ctx sdk.Context, k types.ContractOpsKeeper, p types.InstantiateContractProposal) error {
	if err := p.ValidateBasic(); err != nil {
		return err
	}
	runAsAddr, err := sdk.AccAddressFromBech32(p.RunAs)
	if err != nil {
		return sdkerrors.Wrap(err, "run as address")
	}
	adminAddr, err := sdk.AccAddressFromBech32(p.Admin)
	if err != nil {
		return sdkerrors.Wrap(err, "admin")
	}

	_, data, err := k.Instantiate(ctx, p.CodeID, runAsAddr, adminAddr, p.Msg, p.Label, p.Funds)
	if err != nil {
		return err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeGovContractResult,
		sdk.NewAttribute(types.AttributeKeyResultDataHex, hex.EncodeToString(data)),
	))
	return nil
}

func handleMigrateProposal(ctx sdk.Context, k types.ContractOpsKeeper, p types.MigrateContractProposal) error {
	if err := p.ValidateBasic(); err != nil {
		return err
	}

	contractAddr, err := sdk.AccAddressFromBech32(p.Contract)
	if err != nil {
		return sdkerrors.Wrap(err, "contract")
	}
	runAsAddr, err := sdk.AccAddressFromBech32(p.RunAs)
	if err != nil {
		return sdkerrors.Wrap(err, "run as address")
	}
	data, err := k.Migrate(ctx, contractAddr, runAsAddr, p.CodeID, p.Msg)
	if err != nil {
		return err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeGovContractResult,
		sdk.NewAttribute(types.AttributeKeyResultDataHex, hex.EncodeToString(data)),
	))
	return nil
}

func handleUpdateAdminProposal(ctx sdk.Context, k types.ContractOpsKeeper, p types.UpdateAdminProposal) error {
	if err := p.ValidateBasic(); err != nil {
		return err
	}
	contractAddr, err := sdk.AccAddressFromBech32(p.Contract)
	if err != nil {
		return sdkerrors.Wrap(err, "contract")
	}
	newAdminAddr, err := sdk.AccAddressFromBech32(p.NewAdmin)
	if err != nil {
		return sdkerrors.Wrap(err, "run as address")
	}

	return k.UpdateContractAdmin(ctx, contractAddr, nil, newAdminAddr)
}

func handleClearAdminProposal(ctx sdk.Context, k types.ContractOpsKeeper, p types.ClearAdminProposal) error {
	if err := p.ValidateBasic(); err != nil {
		return err
	}

	contractAddr, err := sdk.AccAddressFromBech32(p.Contract)
	if err != nil {
		return sdkerrors.Wrap(err, "contract")
	}
	if err := k.ClearContractAdmin(ctx, contractAddr, nil); err != nil {
		return err
	}
	return nil
}

func handlePinCodesProposal(ctx sdk.Context, k types.ContractOpsKeeper, p types.PinCodesProposal) error {
	if err := p.ValidateBasic(); err != nil {
		return err
	}
	for _, v := range p.CodeIDs {
		if err := k.PinCode(ctx, v); err != nil {
			return sdkerrors.Wrapf(err, "code id: %d", v)
		}
	}
	return nil
}

func handleUnpinCodesProposal(ctx sdk.Context, k types.ContractOpsKeeper, p types.UnpinCodesProposal) error {
	if err := p.ValidateBasic(); err != nil {
		return err
	}
	for _, v := range p.CodeIDs {
		if err := k.UnpinCode(ctx, v); err != nil {
			return sdkerrors.Wrapf(err, "code id: %d", v)
		}
	}
	return nil
}
