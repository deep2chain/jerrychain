package keeper

import (
	"github.com/deep2chain/jerrychain/x/dex/types"
)

var _ types.QueryServer = Keeper{}
