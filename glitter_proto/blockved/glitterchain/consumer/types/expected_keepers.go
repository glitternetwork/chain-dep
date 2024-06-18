package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here

	GetModuleAddress(name string) sdk.AccAddress
	GetModuleAccount(ctx sdk.Context, moduleName string) types.ModuleAccountI

	SetModuleAccount(sdk.Context, types.ModuleAccountI)
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
}

type IndexKeeper interface {
	CheckDatasetExist(ctx sdk.Context, datasetName string) bool
	AddDatasetStake(ctx sdk.Context, datasetName string, address sdk.AccAddress, amount sdk.Int) bool
}
