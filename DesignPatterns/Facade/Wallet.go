package main

import (
	"fmt"
	"log"
)

type Wallet struct {
	balance int
}

func newWallet() *Wallet {
	return &Wallet{
		balance: 0,
	}
}

func (w *Wallet) creditBalance(amount int) {
	w.balance += amount
	log.Println("wallet balance successfully added!")
	return
}

func (w *Wallet) debitBalance(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("Balance is not sufficient")
	}

	log.Println("balance is updated successfully")
	w.balance -= amount
	return nil
}
