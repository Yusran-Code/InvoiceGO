package model

import (
	"database/sql"
)

// Struct untuk menyimpan data profil pengguna
type AppProfile struct {
	Email           string
	NamaPT          string
	NamaBank        string
	NoRekening      string
	PenanggungJawab string
	Alamat          string
	Kabupaten       string
}
type InvoiceData struct {
	InvoiceNumber string
	InvoiceDate   string
	Periode       string
	QuantityKG    float64
	DisplayQty    float64
	Pokok         float64
	DPP           float64
	PPN           float64
	Total         float64
}

type DataRow struct {
	No          int
	Date        string
	NoSO        string
	NoLO        string
	JumlahTbg   int
	JumlahKg    int
	Tarif       float64
	BiayaAngkut float64
}

func AmbilDataLo(db *sql.DB) ([]DataRow, error) {
	rows, err := db.Query(`
		SELECT 
			no, tanggal,no_so, no_lo, jumlah_tabung, jumlah_kg, tarif, biaya_angkut
		FROM lo_bulanan
		ORDER BY tanggal
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hasil []DataRow
	for rows.Next() {
		var row DataRow
		err := rows.Scan(&row.No, &row.Date, &row.NoSO, &row.NoLO, &row.JumlahTbg, &row.JumlahKg, &row.Tarif, &row.BiayaAngkut)
		if err != nil {
			return nil, err
		}
		hasil = append(hasil, row)
	}
	return hasil, nil
}
