package main

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestConvertDecimalToThaiText(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected string
	}{
		{
			name:     "Required test case 1",
			input:    1234,
			expected: "หนึ่งพันสองร้อยสามสิบสี่บาทถ้วน",
		},
		{
			name:     "Required test case 2",
			input:    33333.75,
			expected: "สามหมื่นสามพันสามร้อยสามสิบสามบาทเจ็ดสิบห้าสตางค์",
		},
		{
			name:     "Zero value",
			input:    0,
			expected: "ศูนย์บาทถ้วน",
		},

		{
			name:     "One baht",
			input:    1,
			expected: "หนึ่งบาทถ้วน",
		},
		{
			name:     "Two baht",
			input:    2,
			expected: "สองบาทถ้วน",
		},
		{
			name:     "Three baht",
			input:    3,
			expected: "สามบาทถ้วน",
		},
		{
			name:     "Four baht",
			input:    4,
			expected: "สี่บาทถ้วน",
		},
		{
			name:     "Five baht",
			input:    5,
			expected: "ห้าบาทถ้วน",
		},
		{
			name:     "Six baht",
			input:    6,
			expected: "หกบาทถ้วน",
		},
		{
			name:     "Seven baht",
			input:    7,
			expected: "เจ็ดบาทถ้วน",
		},
		{
			name:     "Eight baht",
			input:    8,
			expected: "แปดบาทถ้วน",
		},
		{
			name:     "Nine baht",
			input:    9,
			expected: "เก้าบาทถ้วน",
		},
		{
			name:     "Ten baht",
			input:    10,
			expected: "สิบบาทถ้วน",
		},

		{
			name:     "Twenty one",
			input:    21,
			expected: "ยี่สิบเอ็ดบาทถ้วน",
		},
		{
			name:     "One hundred",
			input:    100,
			expected: "หนึ่งร้อยบาทถ้วน",
		},
		{
			name:     "One million",
			input:    1000000,
			expected: "หนึ่งล้านบาทถ้วน",
		},
		{
			name:     "One hundred one with one satang",
			input:    101.01,
			expected: "หนึ่งร้อยเอ็ดบาทหนึ่งสตางค์",
		},
		{
			name:     "One hundred with fifty satang",
			input:    100.50,
			expected: "หนึ่งร้อยบาทห้าสิบสตางค์",
		},
		{
			name:     "Large number with fractional part",
			input:    999999.99,
			expected: "เก้าแสนเก้าหมื่นเก้าพันเก้าร้อยเก้าสิบเก้าบาทเก้าสิบเก้าสตางค์",
		},
		{
			name:     "Small fractional part",
			input:    0.01,
			expected: "ศูนย์บาทหนึ่งสตางค์",
		},
		{
			name:     "Half baht",
			input:    0.50,
			expected: "ศูนย์บาทห้าสิบสตางค์",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := decimal.NewFromFloat(tt.input)
			result := ConvertDecimalToThaiText(input)
			if result != tt.expected {
				t.Errorf("ConvertDecimalToThaiText(%.2f) = %s; expected %s", tt.input, result, tt.expected)
			}
		})
	}
}

func TestConvertIntegerToThai(t *testing.T) {
	tests := []struct {
		name     string
		input    int64
		expected string
	}{
		{
			name:     "Zero",
			input:    0,
			expected: "ศูนย์",
		},
		{
			name:     "One",
			input:    1,
			expected: "หนึ่ง",
		},
		{
			name:     "Two",
			input:    2,
			expected: "สอง",
		},
		{
			name:     "Ten",
			input:    10,
			expected: "สิบ",
		},
		{
			name:     "Eleven",
			input:    11,
			expected: "สิบเอ็ด",
		},
		{
			name:     "Twenty one",
			input:    21,
			expected: "ยี่สิบเอ็ด",
		},
		{
			name:     "One hundred",
			input:    100,
			expected: "หนึ่งร้อย",
		},
		{
			name:     "One thousand",
			input:    1000,
			expected: "หนึ่งพัน",
		},
		{
			name:     "Ten thousand",
			input:    10000,
			expected: "หนึ่งหมื่น",
		},
		{
			name:     "One hundred thousand",
			input:    100000,
			expected: "หนึ่งแสน",
		},

		{
			name:     "One million",
			input:    1000000,
			expected: "หนึ่งล้าน",
		},
		{
			name:     "One trillion",
			input:    1000000000000,
			expected: "หนึ่งล้านล้าน",
		},
		{
			name:     "Complex number",
			input:    1234567,
			expected: "หนึ่งล้านสองแสนสามหมื่นสี่พันห้าร้อยหกสิบเจ็ด",
		},
		{
			name:     "Negative number",
			input:    -1234,
			expected: "ลบหนึ่งพันสองร้อยสามสิบสี่",
		},
		{
			name:     "Hundred thousands - 100000",
			input:    100000,
			expected: "หนึ่งแสน",
		},
		{
			name:     "Hundred thousands - 123400",
			input:    123400,
			expected: "หนึ่งแสนสองหมื่นสามพันสี่ร้อย",
		},
		{
			name:     "Hundred thousands - 500000",
			input:    500000,
			expected: "ห้าแสน",
		},
		{
			name:     "Hundred thousands - 999999",
			input:    999999,
			expected: "เก้าแสนเก้าหมื่นเก้าพันเก้าร้อยเก้าสิบเก้า",
		},
		{
			name:     "Tens of thousands - 12345",
			input:    12345,
			expected: "หนึ่งหมื่นสองพันสามร้อยสี่สิบห้า",
		},
		{
			name:     "Thousands - 1234",
			input:    1234,
			expected: "หนึ่งพันสองร้อยสามสิบสี่",
		},
		{
			name:     "Hundreds - 100",
			input:    100,
			expected: "หนึ่งร้อย",
		},
		{
			name:     "Tens - 21",
			input:    21,
			expected: "ยี่สิบเอ็ด",
		},
		{
			name:     "Special tens - 11",
			input:    11,
			expected: "สิบเอ็ด",
		},
		{
			name:     "Special tens - 10",
			input:    10,
			expected: "สิบ",
		},
		{
			name:     "Special tens - 20",
			input:    20,
			expected: "ยี่สิบ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertIntegerToThai(tt.input)
			if result != tt.expected {
				t.Errorf("convertIntegerToThai(%d) = %s; expected %s", tt.input, result, tt.expected)
			}
		})
	}
}

func BenchmarkConvertDecimalToThaiText(b *testing.B) {
	input := decimal.NewFromFloat(1234567.89)
	for i := 0; i < b.N; i++ {
		ConvertDecimalToThaiText(input)
	}
}

func BenchmarkConvertIntegerToThai(b *testing.B) {
	for i := 0; i < b.N; i++ {
		convertIntegerToThai(1234567)
	}
}
