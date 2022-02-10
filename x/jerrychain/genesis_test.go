package jerrychain_test

import (
	"testing"

	keepertest "github.com/deep2chain/jerrychain/testutil/keeper"
	"github.com/deep2chain/jerrychain/testutil/nullify"
	"github.com/deep2chain/jerrychain/x/jerrychain"
	"github.com/deep2chain/jerrychain/x/jerrychain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.JerrychainKeeper(t)
	jerrychain.InitGenesis(ctx, *k, genesisState)
	got := jerrychain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
