// service/pdf_service.go
package service

import (
	"invoice-go/model"
	"invoice-go/utils"
	"net/http"
	"strconv"

	"github.com/jung-kurt/gofpdf"
)

func GeneratePDFLo(data []model.DataRow, namaPT, bulan string, w http.ResponseWriter) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(0, 10, namaPT)
	pdf.Ln(6)
	pdf.SetFont("Arial", "", 10)
	pdf.Cell(0, 9, "Agen LPG 3 Kg PT Pertamina (Persero)")
	pdf.Ln(6)
	pdf.SetFont("Arial", "B", 11)
	pdf.Cell(0, 9, "Bulan : "+bulan)
	pdf.Ln(12)

	headers := []string{"NO", "Date", "No.SO", "No.LO", "Jum (tbg)", "Jum (Kg)", "TARIF", "BIAYA ANGKUT"}
	widths := []float64{8, 22, 26, 24, 20, 20, 30, 35} // ← Lebarkan tarif & biaya

	pdf.SetFont("Arial", "B", 10)
	for i, h := range headers {
		pdf.CellFormat(widths[i], 6, h, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(-1)

	pdf.SetFont("Arial", "", 10)
	var totalTbg, totalKg int
	for _, row := range data {
		vals := []string{
			strconv.Itoa(row.No),
			row.Date,
			row.NoSO,
			row.NoLO,
			strconv.Itoa(row.JumlahTbg),
			strconv.Itoa(row.JumlahKg),
			"Rp. " + utils.FormatRupiah(row.Tarif),
			"Rp. " + utils.FormatRupiah(row.BiayaAngkut),
		}

		for i, val := range vals {
			pdf.CellFormat(widths[i], 5, val, "1", 0, "C", false, 0, "")
		}
		pdf.Ln(-1)

		totalTbg += row.JumlahTbg
		totalKg += row.JumlahKg
	}

	// Hitung total biaya langsung dari total KG
	totalBiaya := float64(totalKg) * 354.64

	// Baris TOTAL
	pdf.SetFont("Arial", "B", 10)
	for i := 0; i < len(widths); i++ {
		switch i {
		case 3:
			pdf.CellFormat(widths[i], 5, "TOTAL", "1", 0, "R", false, 0, "")
		case 4:
			pdf.CellFormat(widths[i], 5, strconv.Itoa(totalTbg), "1", 0, "C", false, 0, "")
		case 5:
			pdf.CellFormat(widths[i], 5, strconv.Itoa(totalKg), "1", 0, "C", false, 0, "")
		case 6:
			pdf.CellFormat(widths[i], 5, "-", "1", 0, "C", false, 0, "")
		case 7:
			pdf.CellFormat(widths[i], 5, "Rp. "+utils.Formatt(totalBiaya), "1", 0, "C", false, 0, "")
		default:
			pdf.CellFormat(widths[i], 5, "", "1", 0, "C", false, 0, "")
		}
	}
	pdf.Ln(5)

	// PPN dan Grand Total
	ppn := totalBiaya * 0.12
	grandTotal := totalBiaya + ppn

	pdf.SetFont("Arial", "", 10)
	for i := 0; i < len(widths); i++ {
		switch i {
		case 6:
			pdf.CellFormat(widths[i], 5, "PPN 12%", "1", 0, "R", false, 0, "")
		case 7:
			pdf.CellFormat(widths[i], 5, "Rp. "+utils.Formatt(ppn), "1", 0, "R", false, 0, "")
		default:
			pdf.CellFormat(widths[i], 5, " ", "1", 0, "", false, 0, "")
		}
	}
	pdf.Ln(5)

	pdf.SetFont("Arial", "B", 10)
	for i := 0; i < len(widths); i++ {
		switch i {
		case 6:
			pdf.CellFormat(widths[i], 5, "Grand Total", "1", 0, "R", false, 0, "")
		case 7:
			pdf.CellFormat(widths[i], 5, "Rp. "+utils.Formatt(grandTotal), "1", 0, "R", false, 0, "")
		default:
			pdf.CellFormat(widths[i], 5, " ", "1", 0, "", false, 0, "")
		}
	}
	pdf.Ln(30)
	pdf.SetFont("Arial", "", 8)
	pdf.Cell(0, 10, namaPT)

	return pdf.Output(w)
}
