package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

// AddEntry - append entry to the unbonding delegation
func (rCPDT *ReleasingCPDT) AddEntry(creationHeight int64, minTime time.Time, amount sdk.Int) {
	entry := &ReleasingCPDTEntry{
		CreationHeight: creationHeight,
		CompletionTime: minTime,
		Amount:         amount,
	}
	rCPDT.Entries = append(rCPDT.Entries, entry)
}
