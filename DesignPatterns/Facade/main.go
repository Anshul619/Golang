package main

import "log"

func main() {
	walletFacade := newWalletFacade("abc", 1234)

	if err := walletFacade.addMoneyToWallet("abc", 1234, 10); err != nil {
		log.Println("Add Error", err.Error())
	}

	if err := walletFacade.deductMoneyFromWallet("abc", 1234, 5); err != nil {
		log.Println("Debit Error", err.Error())
	}
}
