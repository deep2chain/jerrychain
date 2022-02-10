package keeper_test

import (
	"testing"

	testkeeper "github.com/deep2chain/jerrychain/testutil/keeper"
	"github.com/deep2chain/jerrychain/x/jerrychain/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.JerrychainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
