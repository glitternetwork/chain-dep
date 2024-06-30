package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

const (
	// ModuleName defines the module name

	ModuleName = "consumer"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_consumer"

	AccountCousumerPledgePool = "account_cousumer_pledge_pool"
)

var (
	KeyPrefixConsumer        = []byte{0x10}
	KeyPrefixDatasetStake    = []byte{0x20}
	KeyReleasingCPDT         = []byte{0x30}
	KeyReleasingCPDTQueueKey = []byte{0x41} // prefix for the timestamps in unbonding queue
)

func GetConsumerKey(address []byte) []byte {
	return append(KeyPrefixConsumer, address...)
}

func GetReleasingCPDTKey(address sdk.AccAddress, datasetName string) []byte {
	return append(append(KeyReleasingCPDT, address.Bytes()...), []byte(datasetName)...)
}
func GetReleasingCPDTsKey(address sdk.AccAddress) []byte {
	return append(KeyReleasingCPDT, address.Bytes()...)
}

func GetReleasingCPDTQueueKey(timestamp time.Time) []byte {
	bz := sdk.FormatTimeBytes(timestamp)
	return append(KeyReleasingCPDTQueueKey, bz...)
}
