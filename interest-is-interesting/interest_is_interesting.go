package interest

const (
	negativeBalanceInterest = float32(3.213)
	firstBalanceInterest    = float32(0.5)
	secondBalanceInterest   = float32(1.621)
	thirdBalanceInterest    = float32(2.475)
)

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	if balance < 0 {
		return negativeBalanceInterest
	} else if balance < 1000 {
		return firstBalanceInterest
	} else if balance < 5000 {
		return secondBalanceInterest
	} else {
		return thirdBalanceInterest
	}
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance * float64(InterestRate(balance)) / 100
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	count := 0
	for {
		if balance >= targetBalance {
			return count
		}
		balance += Interest(balance)
		count++
	}
}
