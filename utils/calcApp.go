package utils

import "math"

func HitungTagihan(qtyKg float64) (displayQty, pokok, ppn, total float64) {
	const (
		hargaSatuan  = 560.0
		pengali      = 3.0
		hargaPerGram = 354.64
	)

	// ✅ Hitung displayQty dulu
	displayQty = qtyKg * hargaSatuan * pengali

	// ✅ Karena displayQty sudah gram, pokok tinggal dikali harga per gram
	pokok = displayQty * hargaPerGram

	ppn = math.Round(pokok * 0.12)
	total = pokok + ppn
	return
}
