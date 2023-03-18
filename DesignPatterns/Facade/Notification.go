package main

import "log"

type Notification struct{}

func (n *Notification) sendWalletCreditNotification() {
	log.Println("Sending wallet credit notification")
}

func (n *Notification) sendWalletDebitNotification() {
	log.Println("Sending wallet debit notification")
}
