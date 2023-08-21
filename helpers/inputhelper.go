package helpers

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
	"upcodingactivity/calculator"
)

// ConvertFrequencyInputToNumberOfInterestPaidPerYear This function convert the frequency input to number of interest paid per year based on the rules: 0 = at maturity, 1 = monthly, 2 = quarterly, 3 = annually
func ConvertFrequencyInputToNumberOfInterestPaidPerYear(interestPaidFrequency int) (int, error) {
	numberOfInterestPaidPerYear := 0
	switch interestPaidFrequency {
	case 0:
		numberOfInterestPaidPerYear = 0 // 0 -> by maturity
	case 1:
		numberOfInterestPaidPerYear = 12 // 12 times a year (monthly)
	case 2:
		numberOfInterestPaidPerYear = 4 // 4 times a year (quarterly)
	case 3:
		numberOfInterestPaidPerYear = 1 // 1 time a year (annually)
	default:
		return 0, errors.New("invalid interest paid frequency, should be value of 0,1,2 or 3")
	}
	return numberOfInterestPaidPerYear, nil
}

// TestDataFromCSVFile a helper function to test data from CSV file (refer to test.csv for example)
func TestDataFromCSVFile(cal *calculator.TermDepositCalculator, csvFilePath string, t *testing.T) error {
	csvData, err := os.ReadFile(csvFilePath)
	if err == nil {
		if t == nil {
			log.Println("CSV test file:", csvFilePath)
		}
		r := csv.NewReader(strings.NewReader(string(csvData)))
		meetHeader := false
		recordIndex := 0
		for {
			record, err := r.Read()
			// Skip header
			if !meetHeader {
				meetHeader = true
				continue
			}

			// Increase record index (header not count)
			recordIndex++

			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}

			// Extract and parse data
			depositText := record[0]
			rateText := record[1]
			termText := record[2]
			frequencyText := record[3]
			decimalPlacesText := record[4]
			expectedOutputText := record[5]

			rate, err := strconv.ParseFloat(rateText, 64)
			if err != nil {
				if t != nil {
					t.Fail()
				}
				log.Fatalln("can not passed rate: ", rateText, ", encounter error", err)
			}

			term, err := strconv.Atoi(termText)
			if err != nil {
				if t != nil {
					t.Fail()
				}
				log.Fatalln("can not passed term: ", termText, ", encounter error", err)
			}

			frequency, err := strconv.Atoi(frequencyText)
			if err != nil {
				if t != nil {
					t.Fail()
				}
				log.Fatalln("can not passed frequency: ", frequencyText, ", encounter error", err)
			}

			numberOfInterestPaidPerYear, err := ConvertFrequencyInputToNumberOfInterestPaidPerYear(frequency)
			if err != nil {
				if t != nil {
					t.Fail()
				}
				log.Fatalln(err)
			}

			decimalPlaces, err := strconv.Atoi(decimalPlacesText)
			if err != nil {
				if t != nil {
					t.Fail()
				}
				log.Fatalln("can not passed decimal places: ", decimalPlacesText, ", encounter error", err)
			}

			// Run the calculator to get result
			finalBalance, err := cal.CalculateFinalBalanceWithStringInput(depositText, rate, numberOfInterestPaidPerYear, term)
			if err != nil {
				if t != nil {
					t.Fail()
				}
				log.Fatalln(err)
			} else if finalBalance.StringFixedBank(int32(decimalPlaces)) != expectedOutputText {
				if t != nil {
					t.Fail()
					log.Println("Expected:", expectedOutputText, "Result:", finalBalance.StringFixedBank(int32(decimalPlaces)))
				} else {
					log.Println("FAILED CSV TEST", recordIndex, "\t -> Expected:", expectedOutputText, "\tResult:", finalBalance.StringFixedBank(int32(decimalPlaces)))
				}
			} else {
				if t == nil {
					log.Println("PASSED CSV TEST", recordIndex, "\t -> Expected:", expectedOutputText, "\tResult:", finalBalance.StringFixedBank(int32(decimalPlaces)))
				}
			}
		}
	} else {
		if t != nil {
			t.Fail()
		}
		log.Fatalln("error reading test.csv file", err)
	}
	return err
}
