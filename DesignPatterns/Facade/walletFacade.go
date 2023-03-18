package main

import (
	"log"
)

type WalletFacade struct {
	account      *Account
	wallet       *Wallet
	securityCode *SecurityCode
	notification *Notification
	ledger       *Ledger
}

func newWalletFacade(accountID string, code int) *WalletFacade {
	walletFacade := &WalletFacade{
		account:      newAccount(accountID),
		securityCode: newSecurityCode(code),
		wallet:       newWallet(),
		notification: new(Notification),
		ledger:       new(Ledger),
	}
	log.Println("Account created")
	return walletFacade
}

func (w *WalletFacade) addMoneyToWallet(accountId string, securityCode int, amount int) error {
	log.Println("Starting adding money")
	if err := w.account.checkAccount(accountId); err != nil {
		return err
	}

	if err := w.securityCode.checkCode(securityCode); err != nil {
		return err
	}

	w.wallet.creditBalance(amount)
	w.notification.sendWalletCreditNotification()
	w.ledger.makeEntry(accountId, "credit", amount)
	return nil
}

func (w *WalletFacade) deductMoneyFromWallet(accountId string, securityCode int, amount int) error {
	log.Println("Starting debiting money")
	if err := w.account.checkAccount(accountId); err != nil {
		return err
	}

	if err := w.securityCode.checkCode(securityCode); err != nil {
		return err
	}

	if err := w.wallet.debitBalance(amount); err != nil {
		return err
	}

	w.notification.sendWalletDebitNotification()
	w.ledger.makeEntry(accountId, "credit", amount)
	return nil
}
