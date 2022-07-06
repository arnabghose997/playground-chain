package playgroundchain_test

import (
	"testing"

	keepertest "github.com/arnabghose997/playground-chain/testutil/keeper"
	"github.com/arnabghose997/playground-chain/testutil/nullify"
	"github.com/arnabghose997/playground-chain/x/playgroundchain"
	"github.com/arnabghose997/playground-chain/x/playgroundchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PlaygroundchainKeeper(t)
	playgroundchain.InitGenesis(ctx, *k, genesisState)
	got := playgroundchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
