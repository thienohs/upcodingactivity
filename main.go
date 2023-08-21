package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"upcodingactivity/calculator"
	"upcodingactivity/helpers"

	"github.com/shopspring/decimal"
)

func main() {
	helpText := ". Use --help for more information"

	// Define flags and parse user inputs
	flagLabelDepositAmount := "deposit"
	var depositAmountText string
	flag.StringVar(&depositAmountText, flagLabelDepositAmount, "", "the deposit amount")

	flagLabelInterestRate := "rate"
	var interestRate float64
	flag.Float64Var(&interestRate, flagLabelInterestRate, 0.0, "interest rate (annually)")

	flagLabelInvestmentTerm := "term"
	var investmentTerm int
	flag.IntVar(&investmentTerm, flagLabelInvestmentTerm, 0, "investment term (years)")

	flagLabelInterestPaidFrequency := "frequency"
	var interestPaidFrequency int
	flag.IntVar(&interestPaidFrequency, flagLabelInterestPaidFrequency, 0, "interest paid frequency (0 = at maturity, 1 = monthly, 2 = quarterly, 3 = annually)")

	flagLabelDecimalPlaces := "decimalplaces"
	var decimalPlaces int
	flag.IntVar(&decimalPlaces, flagLabelDecimalPlaces, 0, "decimal places for rounding")

	flagLabelCSVTest := "csvtest"
	var csvTest string
	flag.StringVar(&csvTest, flagLabelCSVTest, "", "the csv file to test (path supported)")
	flag.Parse()

	flagCheck := make(map[string]bool)

	flag.Visit(func(f *flag.Flag) {
		flagCheck[f.Name] = true
	})

	requiredLabels := []string{flagLabelDepositAmount, flagLabelInterestRate, flagLabelInvestmentTerm, flagLabelInterestPaidFrequency}
	for _, label := range requiredLabels {
		_, exist := flagCheck[label]
		if !exist {
			log.Fatalln("please provide:", label, helpText)
		}
	}

	// Validate deposit amount
	if len(depositAmountText) == 0 {
		log.Fatalln("please enter deposit amount. Example: 10000")
	} else {
		_, err := decimal.NewFromString(depositAmountText)
		if err != nil {
			log.Fatalln("invalid deposit amount, encounter error:", err)
		}
	}

	// Validate interestPaidFrequency
	numberOfInterestPaidPerYear, err := helpers.ConvertFrequencyInputToNumberOfInterestPaidPerYear(interestPaidFrequency)
	if err != nil {
		log.Fatalln(err, helpText)
	}

	// Create calculator instance and run with provided inputs
	cal := calculator.NewTermDepositCalculator()
	finalBalance, err := cal.CalculateFinalBalanceWithStringInput(depositAmountText, interestRate, numberOfInterestPaidPerYear, investmentTerm)

	if err == nil {
		fmt.Println("Final balance:", finalBalance.StringFixedBank(int32(decimalPlaces)))
	} else {
		log.Fatalln("error:", err)
	}

	// Validate existence of csv file for testing
	if len(csvTest) > 0 {
		_, err = os.Stat(csvTest)
		if err != nil {
			log.Fatalln("csv file does not exist:", csvTest, helpText)
		} else {
			// Run test data from csv file
			helpers.TestDataFromCSVFile(cal, csvTest, nil)
		}
	}
}
