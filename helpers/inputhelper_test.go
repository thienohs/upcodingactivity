package helpers_test

import (
	"log"
	"testing"
	"upcodingactivity/calculator"
	"upcodingactivity/helpers"
)

// TestConvertByMaturity test convert frequency input "by maturity" to number of interest paid per year
func TestConvertByMaturity(t *testing.T) {
	result, err := helpers.ConvertFrequencyInputToNumberOfInterestPaidPerYear(0)
	expected := 0
	if err != nil {
		t.Fail()
		log.Fatalln(err)
	} else if result != expected {
		t.Fail()
		log.Println("Expected:", expected, "Result:", result)
	}
}

// TestConvertMonthly test convert frequency input "monthly" to number of interest paid per year
func TestConvertMonthly(t *testing.T) {
	result, err := helpers.ConvertFrequencyInputToNumberOfInterestPaidPerYear(1)
	expected := 12
	if err != nil {
		t.Fail()
		log.Fatalln(err)
	} else if result != expected {
		t.Fail()
		log.Println("Expected:", expected, "Result:", result)
	}
}

// TestConvertMonthly test convert frequency input "quarterly" to number of interest paid per year
func TestConvertQuarterly(t *testing.T) {
	result, err := helpers.ConvertFrequencyInputToNumberOfInterestPaidPerYear(2)
	expected := 4
	if err != nil {
		t.Fail()
		log.Fatalln(err)
	} else if result != expected {
		t.Fail()
		log.Println("Expected:", expected, "Result:", result)
	}
}

// TestConvertMonthly test convert frequency input "annually" to number of interest paid per year
func TestConvertAnnually(t *testing.T) {
	result, err := helpers.ConvertFrequencyInputToNumberOfInterestPaidPerYear(3)
	expected := 1
	if err != nil {
		t.Fail()
		log.Fatalln(err)
	} else if result != expected {
		t.Fail()
		log.Println("Expected:", expected, "Result:", result)
	}
}

// TestConvertError test invalid frequency input
func TestConvertError(t *testing.T) {
	_, err := helpers.ConvertFrequencyInputToNumberOfInterestPaidPerYear(4)
	if err == nil {
		t.Fail()
		log.Println("Expected an error but found:", err)
	}
}

// TestUsingTestDataFromCSVFile test using TestDataFromCSVFile for existing valid CSV file
func TestUsingTestDataFromCSVFile(t *testing.T) {
	cal := calculator.NewTermDepositCalculator()
	err := helpers.TestDataFromCSVFile(cal, "../data/test.csv", t)
	if err != nil {
		t.Fail()
		log.Println("ERROR:", err)
	}
}
