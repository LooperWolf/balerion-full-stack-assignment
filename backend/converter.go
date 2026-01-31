package main

import (
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

// ConvertDecimalToThaiText converts a decimal value to Thai text with baht currency suffix
func ConvertDecimalToThaiText(value decimal.Decimal) string {
	// Split the value into integer and fractional parts
	integerPart := value.IntPart()
	fractionalPart := value.Sub(decimal.NewFromInt(integerPart))

	// Convert integer part to Thai text
	integerText := convertIntegerToThai(integerPart)

	// Build the result
	result := integerText + "บาท"

	// Handle fractional part
	if fractionalPart.IsZero() {
		result += "ถ้วน"
	} else {
		// Convert fractional part to satang
		satang := int64(fractionalPart.Mul(decimal.NewFromInt(100)).IntPart())
		satangText := convertIntegerToThai(satang)
		result += satangText + "สตางค์"
	}

	return result
}

// convertIntegerToThai converts an integer to Thai text
func convertIntegerToThai(num int64) string {
	if num == 0 {
		return "ศูนย์"
	}

	if num < 0 {
		return "ลบ" + convertIntegerToThai(-num)
	}

	numStr := strconv.FormatInt(num, 10)
	length := len(numStr)

	if length == 1 {
		return ThaiDigits[num]
	}

	return thaiIntProcessor(numStr)
}

// convertSixDigitNumber converts numbers with 6 or fewer digits to Thai text
func convertSixDigitNumber(numStr string) string {
	length := len(numStr)
	var result strings.Builder
	for i, digit := range numStr {
		digitValue := int(digit - '0')
		position := length - i - 1

		// Skip zero digits except when it's the only digit
		if digitValue == 0 {
			continue
		}

		// Handle special cases for different positions
		switch position {
		case 1: // Tens position
			switch digitValue {
			case 1:
				// For 10-19, we use "สิบ" without the leading "หนึ่ง"
				result.WriteString("สิบ")
			case 2:
				// For 20-29, we use "ยี่สิบ" instead of "สองสิบ"
				result.WriteString("ยี่สิบ")
			default:
				// For 30, 40, 50, etc.
				result.WriteString(ThaiDigits[digitValue])
				result.WriteString("สิบ")
			}
		case 0: // Units position
			if digitValue == 1 && length > 1 {
				// For numbers ending in 1 (except 1 itself), use "เอ็ด"
				result.WriteString("เอ็ด")
			} else {
				result.WriteString(ThaiDigits[digitValue])
			}
		default: // Hundreds, thousands, ten-thousands, hundred-thousands
			result.WriteString(ThaiDigits[digitValue])
			result.WriteString(ThaiPlaces[position])
		}
	}

	return result.String()
}

// thaiIntProcessor handles numbers with more than 6 digits by processing in groups
func thaiIntProcessor(numStr string) string {
	length := len(numStr)

	// Calculate how many complete 6-digit groups we have
	completeGroups := length / 6
	remainder := length % 6

	var result strings.Builder

	// Process the remainder group (leftmost digits that don't make a complete group)
	if remainder > 0 {
		remainderStr := numStr[:remainder]
		remainderText := convertSixDigitNumber(remainderStr)
		if remainderText != "" {
			result.WriteString(remainderText)
			if completeGroups > 0 {
				result.WriteString("ล้าน")
			}
		}
	}

	// Process complete 6-digit groups
	for i := 0; i < completeGroups; i++ {
		start := remainder + (i * 6)
		end := start + 6
		groupStr := numStr[start:end]
		groupText := convertSixDigitNumber(groupStr)

		// Only add group text if it's not all zeros
		if groupText != "" {
			result.WriteString(groupText)
		}

		// Add "ล้าน" for all groups except the last one
		// Also add "ล้าน" if the group is all zeros but not the last group
		if i < completeGroups-1 {
			result.WriteString("ล้าน")
		}
	}

	return result.String()
}
