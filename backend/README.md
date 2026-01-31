# Thai Baht Text Converter

A Go program that converts decimal values to Thai text with "baht" currency suffix.

## Overview

This program converts decimal numbers to their Thai text representation following Thai currency formatting rules. It properly handles both integer amounts (with "ถ้วน" suffix) and amounts with cents (converted to "สตางค์").

## Features

- Converts decimal values to Thai text representation
- Handles both integer and fractional parts
- Properly formats Thai currency with "บาท" (baht) suffix
- Supports special Thai number naming conventions:
  - "หนึ่ง" → "เอ็ด" in units place (except for number 1)
  - "สอง" → "ยี่" in tens place (for numbers > 20)
- Uses the `github.com/shopspring/decimal` package for precise decimal handling

## Requirements

- Go 1.16 or higher
- `github.com/shopspring/decimal` package

## Installation

1. **Install Go** (if not already installed):
   ```bash
   # Ubuntu/Debian
   sudo apt update
   sudo apt install golang-go
   
   # Or download from https://golang.org/dl/
   ```

2. **Clone or navigate to the project directory**:
   ```bash
   cd /home/husky/assignment/bal-assignment/backend
   ```

3. **Install dependencies**:
   ```bash
   go mod tidy
   ```

## Running the Program

### Basic Execution

```bash
go run .
```

This will run the main program with example cases and show the conversion results.

### Running Tests

```bash
# Run all tests with verbose output
go test -v

# Run specific test function
go test -run TestConvertDecimalToThaiText -v

# Run benchmarks
go test -bench=.

# Run tests with coverage
go test -cover

# Run tests with coverage and show which lines are covered
go test -coverprofile=coverage.out && go tool cover -html=coverage.out
```

### Building and Running Binary

```bash
# Build the binary
go build -o thai-baht-converter

# Run the binary
./thai-baht-converter
```

## How It Works

### Function Breakdown

1. **`ConvertDecimalToThaiText(value decimal.Decimal) string`**
   - **Purpose**: Main function that converts decimal values to Thai text
   - **Process**:
     - Separates the input into integer and fractional parts
     - Converts the integer part to Thai text
     - Adds "บาท" suffix
     - For fractional parts:
       - If zero: appends "ถ้วน"
       - If non-zero: converts to "สตางค์"
   - **Return**: Formatted Thai currency string

2. **`convertIntegerToThai(num int64) string`**
   - **Purpose**: Converts integer numbers to Thai text representation
   - **Special Rules**:
     - Number 0: "ศูนย์"
     - Number 1: "หนึ่ง" (unless in compound numbers)
     - "หนึ่ง" → "เอ็ด" when in units place of compound numbers
     - "สอง" → "ยี่" when in tens place (for numbers > 20)
   - **Place Values**: สิบ (10), ร้อย (100), พัน (1,000), หมื่น (10,000), แสน (100,000), ล้าน (1,000,000)

### Example Execution Flow

**Input: 1234.00**
```
1. Separate: Integer = 1234, Fractional = 0.00
2. Convert 1234 to Thai: "หนึ่งพันสองร้อยสามสิบสี่"
3. Add suffix: "บาท" + "ถ้วน"
4. Result: "หนึ่งพันสองร้อยสามสิบสี่บาทถ้วน"
```

**Input: 33333.75**
```
1. Separate: Integer = 33333, Fractional = 0.75
2. Convert 33333 to Thai: "สามหมื่นสามพันสามร้อยสามสิบสาม"
3. Convert fractional 0.75:
   - 0.75 × 100 = 75
   - Convert 75 to Thai: "เจ็ดสิบห้า"
   - Add suffix: "สตางค์"
4. Final result: "สามหมื่นสามพันสามร้อยสามสิบสามบาทเจ็ดสิบห้าสตางค์"
```

## Test Cases

The program includes several test cases:

### Required Test Cases
- **1234**: "หนึ่งพันสองร้อยสามสิบสี่บาทถ้วน"
- **33333.75**: "สามหมื่นสามพันสามร้อยสามสิบสามบาทเจ็ดสิบห้าสตางค์"

### Additional Test Cases
- **0**: "ศูนย์บาทถ้วน"
- **1**: "หนึ่งบาทถ้วน"
- **21**: "ยี่สิบเอ็ดบาทถ้วน"
- **100**: "หนึ่งร้อยบาทถ้วน"
- **1000000**: "หนึ่งล้านบาทถ้วน"
- **101.01**: "หนึ่งร้อยเอ็ดบาทหนึ่งสตางค์"
- **100.50**: "หนึ่งร้อยบาทห้าสิบสตางค์"

## Usage as a Library

The main conversion functions can be easily extracted for use in other projects:

```go
import (
    "github.com/shopspring/decimal"
)

// Convert decimal to Thai text
result := ConvertDecimalToThaiText(decimal.NewFromFloat(1234.50))
// result: "หนึ่งพันสองร้อยสามสิบสี่บาทห้าสิบสตางค์"

// Convert integer to Thai text
text := convertIntegerToThai(1234)
// text: "หนึ่งพันสองร้อยสามสิบสี่"
```

## Project Structure

```
backend/
├── main.go             # Main program with example usage
├── constant.go         # Thai language constants and arrays
├── converter.go        # Core conversion functions
├── converter_test.go   # Comprehensive test suite
├── go.mod              # Go module file
├── go.sum              # Go checksum file
└── README.md           # This documentation
```

## Dependencies

- **github.com/shopspring/decimal**: Provides precise decimal arithmetic for financial calculations

## Error Handling

The program handles the following cases:
- Zero values: Returns "ศูนย์บาทถ้วน"
- Negative values: Prepends "ลบ" to the result
- Large numbers: Supports up to millions (ล้าน)
- Precision: Uses decimal package for accurate fractional calculations

## Contributing

To modify or extend the functionality:
1. Update the conversion logic in `converter.go`
2. Add new test cases to `converter_test.go`
3. Run `go test -v` to verify changes
4. Update this README if functionality changes

### Adding New Test Cases

To add new test cases:
1. Add test cases to the `tests` slice in `TestConvertDecimalToThaiText`
2. Run `go test -run TestConvertDecimalToThaiText -v` to verify
3. Add edge cases to `TestConvertIntegerToThai` if needed

### Performance Testing

Use the benchmark functions to test performance:
```bash
go test -bench=BenchmarkConvertDecimalToThaiText
```

## License

This project is for educational purposes.