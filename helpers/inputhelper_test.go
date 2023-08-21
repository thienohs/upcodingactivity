package helpers_test

import (
	"log"
	"testing"
	"upcodingactivity/calculator"
	"upcodingactivity/helpers"
)

// TestMonthlyInterestPayment_Rate_1_1_Term_3 perform monthly interest payment test
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

func TestConvertError(t *testing.T) {
	_, err := helpers.ConvertFrequencyInputToNumberOfInterestPaidPerYear(4)
	if err == nil {
		t.Fail()
		log.Println("Expected an error but found:", err)
	}
}

func TestUsingTestDataFromCSVFile(t *testing.T) {
	cal := calculator.NewTermDepositCalculator()
	err := helpers.TestDataFromCSVFile(cal, "../data/test.csv", t)
	if err != nil {
		t.Fail()
		log.Println("ERROR:", err)
	}
}
