package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/GlassflowNFT/chtd/x/chtd/types"
    "github.com/GlassflowNFT/chtd/x/chtd/keeper"
    keepertest "github.com/GlassflowNFT/chtd/testutil/keeper"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ChtdKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
