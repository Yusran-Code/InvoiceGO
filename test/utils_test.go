package test

import (
	"fmt"
	"strings"
)

func FormatRupiah(num float64) string {
	s := fmt.Sprintf("%.2f", num)       // pastikan 2 digit di belakang koma
	s = strings.Replace(s, ".", ",", 1) // ubah desimal titik → koma

	// tambahkan titik pemisah ribuan
	parts := strings.Split(s, ",")
	integerPart := parts[0]
	decimalPart := parts[1]

	var result strings.Builder
	n := len(integerPart)
	for i, digit := range integerPart {
		if (n-i)%3 == 0 && i != 0 {
			result.WriteRune('.')
		}
		result.WriteRune(digit)
	}

	return result.String() + "," + decimalPart
}
