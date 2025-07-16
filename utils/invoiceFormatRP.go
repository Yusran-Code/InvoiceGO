package utils

import (
	"strings"

	"github.com/dustin/go-humanize"
)

func Formatt(num float64) string {
	rounded := int64(num)
	s := humanize.Comma(rounded)
	s = strings.Replace(s, ",", ".", -1)
	return s
}

func FormatRupiah(num float64) string {
	s := humanize.CommafWithDigits(num, 2) // pakai 2 desimal
	s = strings.Replace(s, ",", "_", -1)
	s = strings.Replace(s, ".", ",", -1)
	s = strings.Replace(s, "_", ".", -1)
	return s
}
