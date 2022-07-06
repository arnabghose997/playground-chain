package keeper_test

import (
	"testing"

	testkeeper "github.com/arnabghose997/playground-chain/testutil/keeper"
	"github.com/arnabghose997/playground-chain/x/playgroundchain/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.PlaygroundchainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
