package types

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

	AccountCousumerStakePool = "account_cousumer_stake_pool"
)

var (
	KeyPrefixConsumer     = []byte{0x10} // prefix for the historical info
	KeyPrefixDatasetStake = []byte{0x20} // prefix for the historical info
)

func GetConsumerKey(address []byte) []byte {
	return append(KeyPrefixConsumer, address...)
}
