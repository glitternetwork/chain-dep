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

// IsMature - is the current entry mature
func (e *ReleasingCPDTEntry) IsMature(currentTime time.Time) bool {
	return !e.CompletionTime.After(currentTime)
}

func (ubd *ReleasingCPDT) RemoveEntry(i int64) {
	ubd.Entries = append(ubd.Entries[:i], ubd.Entries[i+1:]...)
}
