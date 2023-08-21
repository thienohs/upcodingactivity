package calculator_test

import (
	"log"
	"testing"
	"upcodingactivity/calculator"
	"upcodingactivity/helpers"

	"github.com/shopspring/decimal"
)

// TestMonthlyInterestPayment_Rate_1_1_Term_3 perform monthly interest payment test
func TestMonthlyInterestPayment_Rate_1_1_Term_3(t *testing.T) {
	cal := calculator.NewTermDepositCalculator()
	finalBalance, err := cal.CalculateFinalBalanceWithStringInput("10000", 1.1, 12, 3)
	expectedBalance := "10335"
	if err != nil {
		t.Fail()
		log.Println(err)
	} else if finalBalance.StringFixedBank(0) != expectedBalance {
		t.Fail()
		log.Println("Expected:", expectedBalance, "Result:", finalBalance.StringFixedBank(0))
	}
}

// TestMonthlyInterestPayment_Rate_3_7_Term_5 perform monthly interest payment test
func TestMonthlyInterestPayment_Rate_3_7_Term_5(t *testing.T) {
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

// TestQuarterlyInterestPayment_Rate_1_1_Term_3 perform quarterly interest payment test
func TestQuarterlyInterestPayment_Rate_1_1_Term_3(t *testing.T) {
	cal := calculator.NewTermDepositCalculator()
	finalBalance, err := cal.CalculateFinalBalanceWithStringInput("10000", 1.1, 12/3, 3)
	expectedBalance := "10335"
	if err != nil {
		t.Fail()
		log.Fatalln(err)
	} else if finalBalance.StringFixedBank(0) != expectedBalance {
		t.Fail()
		log.Println("Expected:", expectedBalance, "Result:", finalBalance.StringFixedBank(0))
	}
}

// TestQuarterlyInterestPayment_Rate_3_7_Term_5 perform quarterly interest payment test
func TestQuarterlyInterestPayment_Rate_3_7_Term_5(t *testing.T) {
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

// TestYearlyInterestPayment_Rate_1_1_Term_3 perform yearly interest payment test
func TestYearlyInterestPayment_Rate_1_1_Term_3(t *testing.T) {
	cal := calculator.NewTermDepositCalculator()
	finalBalance, err := cal.CalculateFinalBalanceWithStringInput("10000", 1.1, 1, 3)
	expectedBalance := "10334"
	if err != nil {
		t.Fail()
		log.Fatalln(err)
	} else if finalBalance.StringFixedBank(0) != expectedBalance {
		t.Fail()
		log.Println("Expected:", expectedBalance, "Result:", finalBalance.StringFixedBank(0))
	}
}

// TestYearlyInterestPayment_Rate_3_7_Term_5 perform yearly interest payment test
func TestYearlyInterestPayment_Rate_3_7_Term_5(t *testing.T) {
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

// TestByMaturityInterestPayment_Rate_1_1_Term_3 perform by maturity interest payment test
func TestByMaturityInterestPayment_Rate_1_1_Term_3(t *testing.T) {
	cal := calculator.NewTermDepositCalculator()
	finalBalance, err := cal.CalculateFinalBalanceWithStringInput("10000", 1.1, 0, 3)
	expectedBalance := "10330"
	if err != nil {
		t.Fail()
		log.Fatalln(err)
	} else if finalBalance.StringFixedBank(0) != expectedBalance {
		t.Fail()
		log.Println("Expected:", expectedBalance, "Result:", finalBalance.StringFixedBank(0))
	}
}

// TestByMaturityInterestPayment_Rate_3_7_Term_5 perform by maturity interest payment test
func TestByMaturityInterestPayment_Rate_3_7_Term_5(t *testing.T) {
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

// TestCSV test inputs from CSV file
func TestCSV(t *testing.T) {
	cal := calculator.NewTermDepositCalculator()
	helpers.TestDataFromCSVFile(cal, "../data/test.csv", t)
}
