package atlas_test

import (
	"testing"

	keepertest "github.com/cosmonaut/Atlas/testutil/keeper"
	"github.com/cosmonaut/Atlas/testutil/nullify"
	"github.com/cosmonaut/Atlas/x/atlas"
	"github.com/cosmonaut/Atlas/x/atlas/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AtlasKeeper(t)
	atlas.InitGenesis(ctx, *k, genesisState)
	got := atlas.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
