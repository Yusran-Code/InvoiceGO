package utils

import "math"

func HitungTagihan(qtyKg, dpp float64) (displayQty, pokok, ppn, dppOut, total float64) {
	const (
		hargaSatuan  = 560.0
		pengali      = 3.0
		hargaPerGram = 354.64
	)

	displayQty = qtyKg * hargaSatuan * pengali
	pokok = displayQty * hargaPerGram
	ppn = math.Round(dpp * 0.12)
	dppOut = dpp
	total = pokok + ppn
	return
}
