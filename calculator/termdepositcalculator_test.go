package calculator_test

import (
	"log"
	"testing"
	"upcodingactivity/calculator"
	"upcodingactivity/helpers"

	"github.com/shopspring/decimal"
)

// TestMonthlyInterestPayment perform monthly interest payment test
func TestMonthlyInterestPayment(t *testing.T) {
	depositAmount := decimal.NewFromFloat(10000)
	cal := calculator.NewTermDepositCalculator()
	finalBalance, err := cal.CalculateFinalBalance(&depositAmount, 1.1, 12, 3)
	expectedBalance := "10335"
	if err != nil {
		t.Fail()
		log.Println(err)
	} else if finalBalance.StringFixedBank(0) != expectedBalance {
		t.Fail()
		log.Println("Expected:", expectedBalance, "Result:", finalBalance.StringFixedBank(0))
	}
}

// TestMonthlyInterestPaymentUsingTextInputForDepositAmount perform monthly interest payment test using text input for deposit amount
func TestMonthlyInterestPaymentUsingTextInputForDepositAmount(t *testing.T) {
	cal := calculator.NewTermDepositCalculator()
	finalBalance, err := cal.CalculateFinalBalanceWithStringInput("10000", 3.7, 12, 5)
	expectedBalance := "12029"
	if err != nil {
		t.Fail()
		log.Println(err)
	} else if finalBalance.StringFixedBank(0) != expectedBalance {
		t.Fail()
		log.Println("Expected:", expectedBalance, "Result:", finalBalance.StringFixedBank(0))
	}
}

// TestQuarterlyInterestPayment perform quarterly interest payment test
func TestQuarterlyInterestPayment(t *testing.T) {
	depositAmount := decimal.NewFromFloat(10000)
	cal := calculator.NewTermDepositCalculator()
	finalBalance, err := cal.CalculateFinalBalance(&depositAmount, 1.1, 12/3, 3)
	expectedBalance := "10335"
	if err != nil {
		t.Fail()
		log.Fatalln(err)
	} else if finalBalance.StringFixedBank(0) != expectedBalance {
		t.Fail()
		log.Println("Expected:", expectedBalance, "Result:", finalBalance.StringFixedBank(0))
	}
}

// TestQuarterlyInterestPaymentUsingTextInputForDepositAmount perform quarterly interest payment test using text input for deposit amount
func TestQuarterlyInterestPaymentUsingTextInputForDepositAmount(t *testing.T) {
	cal := calculator.NewTermDepositCalculator()
	finalBalance, err := cal.CalculateFinalBalanceWithStringInput("10000", 3.7, 12/3, 5)
	expectedBalance := "12022"
	if err != nil {
		t.Fail()
		log.Fatalln(err)
	} else if finalBalance.StringFixedBank(0) != expectedBalance {
		t.Fail()
		log.Println("Expected:", expectedBalance, "Result:", finalBalance.StringFixedBank(0))
	}
}

// TestYearlyInterestPayment perform yearly interest payment test
func TestYearlyInterestPayment(t *testing.T) {
	depositAmount := decimal.NewFromFloat(10000)
	cal := calculator.NewTermDepositCalculator()
	finalBalance, err := cal.CalculateFinalBalance(&depositAmount, 1.1, 1, 3)
	expectedBalance := "10334"
	if err != nil {
		t.Fail()
		log.Fatalln(err)
	} else if finalBalance.StringFixedBank(0) != expectedBalance {
		t.Fail()
		log.Println("Expected:", expectedBalance, "Result:", finalBalance.StringFixedBank(0))
	}
}

// TestYearlyInterestPaymentUsingTextInputForDepositAmount perform yearly interest payment test using text input for deposit amount
func TestYearlyInterestPaymentUsingTextInputForDepositAmount(t *testing.T) {
	cal := calculator.NewTermDepositCalculator()
	finalBalance, err := cal.CalculateFinalBalanceWithStringInput("10000", 3.7, 1, 5)
	expectedBalance := "11992"
	if err != nil {
		t.Fail()
		log.Fatalln(err)
	} else if finalBalance.StringFixedBank(0) != expectedBalance {
		t.Fail()
		log.Println("Expected:", expectedBalance, "Result:", finalBalance.StringFixedBank(0))
	}
}

// TestByMaturityInterestPayment perform by maturity interest payment test
func TestByMaturityInterestPayment(t *testing.T) {
	depositAmount := decimal.NewFromFloat(10000)
	cal := calculator.NewTermDepositCalculator()
	finalBalance, err := cal.CalculateFinalBalance(&depositAmount, 1.1, 0, 3)
	expectedBalance := "10330"
	if err != nil {
		t.Fail()
		log.Fatalln(err)
	} else if finalBalance.StringFixedBank(0) != expectedBalance {
		t.Fail()
		log.Println("Expected:", expectedBalance, "Result:", finalBalance.StringFixedBank(0))
	}
}

// TestByMaturityInterestPaymentUsingTextInputForDepositAmount perform by maturity interest payment test using text input for deposit amount
func TestByMaturityInterestPaymentUsingTextInputForDepositAmount(t *testing.T) {
	cal := calculator.NewTermDepositCalculator()
	finalBalance, err := cal.CalculateFinalBalanceWithStringInput("10000", 3.7, 0, 5)
	expectedBalance := "11850"
	if err != nil {
		t.Fail()
		log.Fatalln(err)
	} else if finalBalance.StringFixedBank(0) != expectedBalance {
		t.Fail()
		log.Println("Expected:", expectedBalance, "Result:", finalBalance.StringFixedBank(0))
	}
}

// TestLargeInput perform test with large input
func TestLargeInput(t *testing.T) {
	cal := calculator.NewTermDepositCalculator()
	finalBalance, err := cal.CalculateFinalBalanceWithStringInput("1000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", 1.1, 12, 3)
	expectedBalance := "1033534916454078225679782255448432722745079551427440206564576599859159158471212532887402583138973419665405227371992"
	if err != nil {
		t.Fail()
		log.Fatalln(err)
	} else if finalBalance.StringFixedBank(0) != expectedBalance {
		t.Fail()
		log.Println("Expected:", expectedBalance, "Result:", finalBalance.StringFixedBank(0))
	}
}

// TestPrecision test precision
func TestPrecision(t *testing.T) {
	cal := calculator.NewTermDepositCalculator()
	depositAmount := decimal.NewFromFloat(10000)
	finalBalance, err := cal.CalculateFinalBalance(&depositAmount, 1.1, 1, 3)
	expectedBalance := "10333.6433100000000000000000000000000000000000000000000000000000000000"
	precision := int32(64)
	if err != nil {
		t.Fail()
		log.Fatalln(err)
	} else if finalBalance.StringFixedBank(precision) != expectedBalance {
		t.Fail()
		log.Println("Expected:", expectedBalance, "Result:", finalBalance.StringFixedBank(precision))
	}
}

// TestOutOfRangeNumberOfInterestPaidPerYear test when numberOfInterestPaidPerYear is out of range (current valid range is 0-52)
func TestOutOfRangeNumberOfInterestPaidPerYear(t *testing.T) {
	depositAmount := decimal.NewFromFloat(10000)
	cal := calculator.NewTermDepositCalculator()
	_, err := cal.CalculateFinalBalance(&depositAmount, 1.1, 53, 3)
	if err == nil {
		t.Fail()
		log.Println("Expected an error but found: ", err)
	}

	_, err = cal.CalculateFinalBalance(&depositAmount, 1.1, -1, 3)
	if err == nil {
		t.Fail()
		log.Println("Expected an error but found: ", err)
	}
}

// TestCSV test inputs from CSV file
func TestCSV(t *testing.T) {
	cal := calculator.NewTermDepositCalculator()
	helpers.TestDataFromCSVFile(cal, "../data/test.csv", t)
}
