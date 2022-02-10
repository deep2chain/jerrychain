package keeper

import (
	"github.com/deep2chain/jerrychain/x/jerrychain/types"
)

var _ types.QueryServer = Keeper{}
