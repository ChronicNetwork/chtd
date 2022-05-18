package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/ChronicNetwork/chtd/testutil/keeper"
	"github.com/ChronicNetwork/chtd/x/chtd/keeper"
	"github.com/ChronicNetwork/chtd/x/chtd/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ChtdKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
