package utils

import (
	"invoice-go/model"

	"github.com/dustin/go-humanize"
	"github.com/jung-kurt/gofpdf"
)

func GeneratePDFInvoice(profile model.AppProfile, data model.InvoiceData) *gofpdf.Fpdf {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(20, 20, 20)
	pdf.AddPage()

	// Header dan alamat
	pdf.SetFont("Arial", "B", 14)
	pdf.CellFormat(0, 7, profile.NamaPT, "", 1, "C", false, 0, "")
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(0, 5, "Agen LPG PSO", "", 1, "C", false, 0, "")
	pdf.CellFormat(0, 5, profile.Alamat+" , "+profile.Kabupaten, "", 1, "C", false, 0, "")
	pdf.Ln(5)

	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(0, 10, "INVOICE", "", 1, "C", false, 0, "")
	pdf.Ln(10)

	// Kepada
	pdf.SetFont("Arial", "", 9)
	pdf.Cell(0, 6, "Kepada : PT.Pertamina Patra Niaga")
	pdf.Ln(6)
	pdf.Cell(0, 6, "Alamat : Gedung Wisma Tugu II Lt.2, Jl. HR Rasuna Said KAV C7-9 Setiabudi, Jakarta 12920")
	pdf.Ln(10)

	// Nomor dan tanggal invoice
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(95, 6, "Tanggal: "+data.InvoiceDate, "", 0, "L", false, 0, "")
	pdf.SetX(115)
	pdf.CellFormat(0, 6, "No. Invoice: "+data.InvoiceNumber, "", 1, "R", false, 0, "")
	pdf.Ln(5)

	// Tabel
	pdf.SetFont("Arial", "B", 11)
	pdf.SetFillColor(240, 240, 240)
	pdf.CellFormat(90, 8, "Deskripsi", "1", 0, "C", true, 0, "")
	pdf.CellFormat(90, 8, "Nilai", "1", 1, "C", true, 0, "")

	pdf.SetFont("Arial", "", 11)
	pdf.CellFormat(90, 8, "Tagihan Transport Fee Periode "+data.Periode, "1", 0, "L", false, 0, "")
	pdf.CellFormat(90, 8, "", "1", 1, "R", false, 0, "")
	pdf.CellFormat(90, 8, "Quantity/Kg", "1", 0, "L", false, 0, "")
	pdf.CellFormat(90, 8, humanize.Comma(int64(data.DisplayQty)), "1", 1, "R", false, 0, "")
	pdf.CellFormat(90, 8, "DPP", "1", 0, "L", false, 0, "")
	pdf.CellFormat(90, 8, humanize.Comma(int64(data.DPP)), "1", 1, "R", false, 0, "")
	pdf.CellFormat(90, 8, "Pokok", "1", 0, "L", false, 0, "")
	pdf.CellFormat(90, 8, humanize.Comma(int64(data.Pokok)), "1", 1, "R", false, 0, "")
	pdf.CellFormat(90, 8, "PPN 12%", "1", 0, "L", false, 0, "")
	pdf.CellFormat(90, 8, humanize.Comma(int64(data.PPN)), "1", 1, "R", false, 0, "")

	pdf.SetFont("Arial", "B", 11)
	pdf.CellFormat(90, 8, "Total", "1", 0, "L", false, 0, "")
	pdf.CellFormat(90, 8, "Rp. "+humanize.Comma(int64(data.Total)), "1", 1, "R", false, 0, "")

	pdf.Ln(8)
	pdf.SetFont("Arial", "I", 10)
	pdf.MultiCell(0, 6, "Terbilang: "+Terbilang(int64(data.Total)), "", "", false)

	pdf.SetFont("Arial", "", 10)
	pdf.Ln(15)
	pdf.Cell(0, 6, "Bank: "+profile.NamaBank+" - "+profile.NoRekening)
	pdf.Ln(6)
	pdf.Cell(0, 6, "a.n "+profile.NamaPT)

	pdf.Ln(35)
	pdf.CellFormat(0, 6, profile.PenanggungJawab, "", 1, "R", false, 0, "")
	return pdf
}
