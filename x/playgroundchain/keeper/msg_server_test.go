package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/arnabghose997/playground-chain/testutil/keeper"
	"github.com/arnabghose997/playground-chain/x/playgroundchain/keeper"
	"github.com/arnabghose997/playground-chain/x/playgroundchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.PlaygroundchainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
