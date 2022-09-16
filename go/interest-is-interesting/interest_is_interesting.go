package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	var interest float32
	if balance < 0 {
		interest = 3.213
	} else if balance >= 0 && balance < 1000 {
		interest = 0.5
	} else if balance >= 1000 && balance < 5000 {
		interest = 1.621
	} else if balance >= 5000 {
		interest = 2.475
	}
	return interest
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance * float64(InterestRate(balance)/100)
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	year := 0
	for balance < targetBalance {
		year++
		balance = AnnualBalanceUpdate(balance)
	}
	return year
}
