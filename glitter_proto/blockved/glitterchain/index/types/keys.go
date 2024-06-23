package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	"time"
)

const (
	// ModuleName defines the module name
	ModuleName = "index"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_index"

	AccountDatasePurchaseFeePool = "account_datase_purchase_fee_pool"
)

var (
	KeyDataset               = []byte{0x10}
	KeyTable                 = []byte{0x20}
	KeyConsumerPledge        = []byte{0x30}
	KeyDatasetExpiratioQueue = []byte{0x40}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func GetDatasetKey(datasetName string) []byte {
	return append(KeyDataset, []byte(datasetName)...)
}

func GetTableKey(datasetName string, tableName string) []byte {
	return append(KeyTable, []byte(datasetName+":"+tableName)...)
}

func GetConsumerPledgeKey(datasetName string, consumerAddress sdk.AccAddress) []byte {
	return append(GetConsumerPledgesKey(datasetName), address.MustLengthPrefix(consumerAddress)...)
}

func GetConsumerPledgesKey(datasetName string) []byte {
	return append(KeyTable, []byte(datasetName)...)
}

func GetDatasetExpirationKey(timestamp time.Time) []byte {
	bz := sdk.FormatTimeBytes(timestamp)
	return append(KeyDatasetExpiratioQueue, bz...)
}
