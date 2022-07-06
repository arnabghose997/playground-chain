package keeper

import (
	"github.com/arnabghose997/playground-chain/x/playgroundchain/types"
)

var _ types.QueryServer = Keeper{}
