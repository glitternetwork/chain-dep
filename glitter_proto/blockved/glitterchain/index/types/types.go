package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

const (
	GasAmountESCreateMapping     = sdk.Gas(1)
	GasDescriptorESCreateMapping = "create mapping"

	GasAmountESInsertDoc   = sdk.Gas(1)
	GasDescriptorInsertDoc = "insert doc"

	GasAmountCreateTablePerField = sdk.Gas(1000)
	GasAmountInsertTablePerByte  = sdk.Gas(1)

	GasDescriptorExecSql = "exec sql"
)

var GlobalAccAddrDatasetStake = sdk.AccAddress(crypto.AddressHash([]byte("GlobalAccAddrDatasetStake")))
