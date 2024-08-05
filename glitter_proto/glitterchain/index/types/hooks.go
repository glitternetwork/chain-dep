package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// combine multiple staking hooks, all hook functions are run in array sequence
type MultiIndexHooks []IndexHooks

func NewMultiIndexHooks(hooks ...IndexHooks) MultiIndexHooks {
	return hooks
}

func (h MultiIndexHooks) AfterDatasetExpired(ctx sdk.Context, datasetName string, address sdk.AccAddress) {
	for i := range h {
		h[i].AfterDatasetExpired(ctx, datasetName, address)
	}
}
