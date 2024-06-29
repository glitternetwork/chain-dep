package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

// ReleasingCPDT
func (rCPDT *ReleasingCPDT) AddEntry(creationHeight int64, minTime time.Time, amount sdk.Int) {
	entry := &ReleasingCPDTEntry{
		CreationHeight: creationHeight,
		CompletionTime: minTime,
		Amount:         amount,
	}
	rCPDT.Entries = append(rCPDT.Entries, entry)
}

func (rCPDT *ReleasingCPDT) RemoveEntryByIndex(i int64) {
	rCPDT.Entries = append(rCPDT.Entries[:i], rCPDT.Entries[i+1:]...)
}

// ReleasingCPDTEntry
// IsMature - is the current entry mature
func (e *ReleasingCPDTEntry) IsMature(currentTime time.Time) bool {
	return !e.CompletionTime.After(currentTime)
}

// consumer
func (c *Consumer) RemoveCPDTByIndex(i int64) {
	c.CPDTs = append(c.CPDTs[:i], c.CPDTs[i+1:]...)
}
