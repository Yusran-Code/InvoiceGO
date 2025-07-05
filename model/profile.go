package model

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
