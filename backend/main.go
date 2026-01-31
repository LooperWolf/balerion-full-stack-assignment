package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	// Example usage of the Thai Baht text converter
	fmt.Println("Thai Baht Text Converter Examples")
	fmt.Println("================================")

	// Required test cases from the specification
	examples := []struct {
		value       float64
		description string
	}{
		{1234, "หนึ่งพันสองร้อยสามสิบสี่บาทถ้วน"},
		{33333.75, "สามหมื่นสามพันสามร้อยสามสิบสามบาทเจ็ดสิบห้าสตางค์"},
		{0, "ศูนย์บาทถ้วน"},
		{1, "หนึ่งบาทถ้วน"},
		{21, "ยี่สิบเอ็ดบาทถ้วน"},
		{100, "หนึ่งร้อยบาทถ้วน"},
		{1000000, "หนึ่งล้านบาทถ้วน"},
		{100000000, "หนึ่งร้อยล้านบาทถ้วน"},
		{101.01, "หนึ่งร้อยเอ็ดบาทหนึ่งสตางค์"},
	}

	for _, example := range examples {
		input := decimal.NewFromFloat(example.value)
		result := ConvertDecimalToThaiText(input)
		fmt.Printf("Input: %.2f\n", example.value)
		fmt.Printf("Output: %s\n", result)
		fmt.Printf("Expected: %s\n", example.description)
		fmt.Printf("Match: %t\n\n", result == example.description)
	}

	fmt.Println("To run comprehensive tests, use: go test -v")
	fmt.Println("To run benchmarks, use: go test -bench=.")
}
