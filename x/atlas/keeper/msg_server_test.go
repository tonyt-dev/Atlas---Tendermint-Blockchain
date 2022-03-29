package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/cosmonaut/Atlas/testutil/keeper"
	"github.com/cosmonaut/Atlas/x/atlas/keeper"
	"github.com/cosmonaut/Atlas/x/atlas/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.AtlasKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
