package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

// AddEntry - append entry to the unbonding delegation
func (ubds *UnbondingStakePair) AddEntry(creationHeight int64, minTime time.Time, amount sdk.Int) {
	entry := &UnbondingStakeEntry{
		CreationHeight: creationHeight,
		CompletionTime: minTime,
		Amount:         amount,
	}
	ubds.Entries = append(ubds.Entries, entry)
}
