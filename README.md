
# Project description
Simple term deposit calculator that takes as inputs:
- Start deposit amount (e.g. $10,000)
- Interest rate (e.g. 1.10%)
- Investment term (e.g. 3 years)
- Interest paid (Monthly, Quarterly, Annually, At Maturity)

And produces as output:
- Final balance (e.g. $10,330 on the above inputs, interest paid At Maturity)

# Calculator rules

## Compound interest formula
Compound interest, can be calculated using the formula FV = P*(1+R/N)^(N*T), where:
- FV is the future value of the loan or investment
- P is the initial principal amount (depositAmount)
- R is the annual interest rate (interestRateAnnually)
- N represents the number of times interest is compounded per year
- T represents time in years (investmentTermInYears)

Credit: https://www.realized1031.com/glossary/compound-interest

# Consideration:
- Decimal to be used for calculating currency value
- Precision when rounding

# Usage

## Installation
- Download and install Go: https://go.dev/doc/install
- Version used: Go 1.19

## Note: 
- Please run the commands at the project folder
- Required flags:
    - -deposit (Deposit amount)
    - -rate (interest rate, annually)
    - -term (investment term, years)
    - -frequency (interest paid frequency)
- Optional flags: 
    - -decimalplaces (Number of decimal places for rounding)
    - -csvtest (CSV file path to test)
- Accepted values for interest paid frequency are: 0,1,2 or 3 (which 0 = at maturity, 1 = monthly, 2 = quarterly, 3 = annually)
- Deposit amount can support large number (maximum of 2^31 digits after the decimal point)
- Interest rate is presented using float64
- Investment term is presented using int (int is a signed integer type that is at least 32 bits in size)
- Decimal places is presented using int
- CSV test file path can be provided following the example format in data/test.csv to perform multiple tests

## Normal usage with support for deposit amount, interest rate, investment term and interest paid frequency
```
go run . -deposit 10000  -rate 1.1 -term 30 -frequency 3
```

## With decimal places
```
go run . -deposit 10000  -rate 1.1 -term 30 -frequency 3 -decimalplaces=2
```

## Test multiple inputs using csv file
```
go run . -deposit 10000  -rate 1.1 -term 30 -frequency 3 -csvtest data/test.csv
```

## Help
```
go run . --help
```

# Testing

## Run all available test cases
```
go test ./...
```

## Run all available test cases (verbose)
```
go test ./... -v
```

# Build instruction

## Current platform
```
go build -o bin/upcodingactivity
```

## MacOS (64-bit)
```
GOOS=darwin GOARCH=amd64 go build -o bin/mac/upcodingactivity
```

## MacOS (M1, 64-bit)
```
GOOS=darwin GOARCH=arm64 go build -o bin/mac/upcodingactivity
```

## Windows (64-bit)
```
GOOS=windows GOARCH=amd64 go build -o bin/windows/upcodingactivity.exe
```

## Linux (64-bit)
```
GOOS=linux GOARCH=amd64  go build -o bin/linux/upcodingactivity
```

## Binany can be used to execute without using go commands. Examples:
```
./bin/mac/upcodingactivity -deposit 10000  -rate 1.1 -term 30 -frequency 3
```
```
./bin/mac/upcodingactivity -deposit 10000  -rate 1.1 -term 30 -frequency 3 -decimalplaces=2
```

# Opensource licenses and credits
- https://github.com/shopspring/decimal (MIT)

# Deployment
N/A