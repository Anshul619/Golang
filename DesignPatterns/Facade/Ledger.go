package main

import "fmt"

type Ledger struct {
}

func (s *Ledger) makeEntry(accountId, txnType string, amount int) {
	fmt.Println("Make ledger entry for accountId %s with txnType %s for amount %d\n", accountId, txnType, amount)
	return
}
