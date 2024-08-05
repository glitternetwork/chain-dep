package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

var GlobalAccAddrDatasetStake = sdk.AccAddress(crypto.AddressHash([]byte("GlobalAccAddrDatasetStake")))

func MustUnmarshalDataset(cdc codec.BinaryCodec, value []byte) (r Dataset) {
	r = Dataset{}
	err := cdc.Unmarshal(value, &r)
	if err != nil {
		panic(err)
	}
	return r
}

func MustUnmarshalCPDT(cdc codec.BinaryCodec, value []byte) (r CPDT) {
	r = CPDT{}
	err := cdc.Unmarshal(value, &r)
	if err != nil {
		panic(err)
	}
	return r
}

func MustUnmarshalDatasetExpiration(cdc codec.BinaryCodec, value []byte) (r DatasetExpiration) {
	r = DatasetExpiration{}
	err := cdc.Unmarshal(value, &r)
	if err != nil {
		panic(err)
	}
	return r
}
