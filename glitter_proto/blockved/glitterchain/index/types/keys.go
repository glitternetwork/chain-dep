package types

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

	SchemaKey = "schema_"

	DocKey = "doc_"
)

var (
	HostKeyByDB      = []byte{0x50} // prefix for the historical info
	KeyPrefixDataset = []byte{0x60} // prefix for the historical info
	KeyPrefixTable   = []byte{0x70} // prefix for the historical info
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func GetHostKeyByDB(db string) []byte {
	return append(HostKeyByDB, []byte(db)...)
}

func GetDatasetKey(datasetName string) []byte {
	return append(KeyPrefixTable, []byte(datasetName)...)
}

func GetTableKey(datasetName string, tableName string) []byte {
	return append(KeyPrefixTable, []byte(datasetName+":"+tableName)...)
}
