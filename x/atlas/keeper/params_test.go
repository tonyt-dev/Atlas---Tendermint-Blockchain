package keeper_test

import (
	"testing"

	testkeeper "github.com/cosmonaut/Atlas/testutil/keeper"
	"github.com/cosmonaut/Atlas/x/atlas/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.AtlasKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
