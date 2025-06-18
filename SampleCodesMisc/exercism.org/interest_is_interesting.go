package main

import "log"

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	switch {
	    case balance < 0:
	        return 3.213
	    case balance >= 0 && balance < 1000:
	        return 0.5
	    case balance >= 1000 && balance < 5000:
            return 1.621
	}

	return 2.475
}

func Interest(balance float64) float64 {
    return (balance*float64(InterestRate(balance)))/100
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance+Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {

    term := 0

    for targetBalance > balance {
        balance = AnnualBalanceUpdate(balance)
        term++
    }

	return term
}

func main() {
    log.Println(InterestRate(200.75))
    log.Println(Interest(200.75))
    log.Println(AnnualBalanceUpdate(200.75))
    log.Println(YearsBeforeDesiredBalance(200.75, 214.88))
    log.Println(YearsBeforeDesiredBalance(100.0, 125.80))
    log.Println(YearsBeforeDesiredBalance(2345.67, 12345.6789))
}

