package calculator

import (
	"errors"

	"github.com/shopspring/decimal"
)

type TermDepositCalculator struct {
}

// NewTermDepositCalculator create a new term deposit calculator instance
func NewTermDepositCalculator() *TermDepositCalculator {
	return &TermDepositCalculator{}
}

// CalculateFinalBalance this function will calculate final balance with inputs:
// depositAmount(Arbitrary-precision fixed-point decimal number) the original deposit amount (principal amount),
// interestRateAnnually interest rate (annual),
// interestPaidFrequencyInMonth interest paid out frequency (in number of months),
// investmentTermInYears investment term (in number of years)
//
// Return: the final balance and error (if any)
func (c *TermDepositCalculator) CalculateFinalBalance(depositAmount *decimal.Decimal, interestRateAnnually float64, numberOfInterestPaidPerYear int, investmentTermInYears int) (*decimal.Decimal, error) {
	// Validation
	// 0 <= interestPaidFrequencyInMonth <= 52, which interestPaidFrequencyInMonth = 0 -> By maturity
	if numberOfInterestPaidPerYear < 0 || numberOfInterestPaidPerYear > 52 {
		return nil, errors.New("numberOfInterestPaidPerYear is out of range, valid range is 0 - 52")
	}

	// This will calculate the fixed annual interest earned
	fixedAnnualInterestEarned := decimal.NewFromFloat(float64(interestRateAnnually))
	fixedAnnualInterestEarned = fixedAnnualInterestEarned.Mul(*depositAmount)
	fixedAnnualInterestEarned = fixedAnnualInterestEarned.Div(decimal.NewFromFloat((100)))

	// Use the fixedAnnualInterestEarned to go through re-investment process

	// Compound interest, can be calculated using the formula FV = P*(1+R/N)^(N*T),
	// where FV is the future value of the loan or investment,
	// P is the initial principal amount (depositAmount),
	// R is the annual interest rate (interestRateAnnually),
	// N represents the number of times interest is compounded per year (numberOfInterestPaidPerYear, if numberOfInterestPaidPerYear = 0 -> By maturity),
	// T represents time in years. (investmentTermInYears)
	// Credit: https://www.realized1031.com/glossary/compound-interest
	finalBalance := decimal.NewFromFloat(0)

	// if numberOfInterestPaidPerYear = 0 -> By maturity
	if numberOfInterestPaidPerYear == 0 {
		// Final balance = depositAmount + fixedAnnualInterestEarned * investmentTermInYears
		finalBalance = (*depositAmount).Add(fixedAnnualInterestEarned.Mul(decimal.NewFromFloat(float64(investmentTermInYears))))
	} else {
		noOfTimesInterestIsCompoundedPerYear := decimal.NewFromFloat(float64(numberOfInterestPaidPerYear))
		// Part 1 = (1+R/N)
		part1 := decimal.NewFromFloat(interestRateAnnually / 100)
		part1 = part1.Div(noOfTimesInterestIsCompoundedPerYear)
		part1 = part1.Add(decimal.NewFromFloat(1.0))

		// Part 2 = (N*T)
		part2 := noOfTimesInterestIsCompoundedPerYear.Mul(decimal.NewFromFloat(float64(investmentTermInYears)))

		// Final balance = depositAmount * (part1 ^ part2)
		finalBalance = part1.Pow(part2).Mul(*depositAmount)
	}

	return &finalBalance, nil
}

// calculateFinalBalanceWithStringInput this function will calculate final balance using CalculateFinalBalance with inputs:
// depositAmount(string) the original deposit amount (principal amount),
// interestRateAnnually interest rate (annual),
// interestPaidFrequencyInMonth interest paid out frequency (in number of months),
// investmentTermInYears investment term (in number of years)
//
// Return: the final balance and error (if any)
func (c *TermDepositCalculator) CalculateFinalBalanceWithStringInput(depositAmountText string, interestRateAnnually float64, numberOfInterestPaidPerYear int, investmentTermInYears int) (*decimal.Decimal, error) {
	// Parse deposit decimal amount from input text
	depositAmount, err := decimal.NewFromString(depositAmountText)
	if err == nil {
		return c.CalculateFinalBalance(&depositAmount, interestRateAnnually, numberOfInterestPaidPerYear, investmentTermInYears)
	}
	return nil, err
}
