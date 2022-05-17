package chtd_test

import (
	"testing"

	keepertest "github.com/GlassflowNFT/chtd/testutil/keeper"
	"github.com/GlassflowNFT/chtd/testutil/nullify"
	"github.com/GlassflowNFT/chtd/x/chtd"
	"github.com/GlassflowNFT/chtd/x/chtd/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:	types.DefaultParams(),
		
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ChtdKeeper(t)
	chtd.InitGenesis(ctx, *k, genesisState)
	got := chtd.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	

	// this line is used by starport scaffolding # genesis/test/assert
}
