package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/deep2chain/jerrychain/testutil/keeper"
	"github.com/deep2chain/jerrychain/x/jerrychain/keeper"
	"github.com/deep2chain/jerrychain/x/jerrychain/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.JerrychainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
