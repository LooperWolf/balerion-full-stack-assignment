package main

// ThaiDigits represents Thai number words from 0-9

var ThaiDigits = []string{
	"",      // 0
	"หนึ่ง", // 1
	"สอง",   // 2
	"สาม",   // 3
	"สี่",   // 4
	"ห้า",   // 5
	"หก",    // 6
	"เจ็ด",  // 7
	"แปด",   // 8
	"เก้า",  // 9
}

// ThaiPlaces represents Thai place values
var ThaiPlaces = []string{"", "สิบ", "ร้อย", "พัน", "หมื่น", "แสน", "ล้าน"}

// ThaiExtendedPlaces represents extended place values for very large numbers
var ThaiExtendedPlaces = map[int]string{
	7:  "สิบล้าน",
	8:  "พันล้าน",
	9:  "หมื่นล้าน",
	10: "แสนล้าน",
	11: "ล้านล้าน",
	12: "สิบล้านล้าน",
}
