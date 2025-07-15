package service

import (
	"dasar-go/model"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/jung-kurt/gofpdf"
)

func handleGeneratePDF(w http.ResponseWriter, r *http.Request, isDownload bool) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Gagal parsing form: "+err.Error(), http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Gagal membaca file: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	data, err := ParseExcelToDataRows(file)
	if err != nil || len(data) == 0 {
		http.Error(w, "Gagal parsing data Excel", http.StatusBadRequest)
		return
	}

	namaPT := r.FormValue("namapt")
	bulan := r.FormValue("bulan")
	if bulan == "" {
		bulan = time.Now().Format("January 2006")
	}

	w.Header().Set("Content-Type", "application/pdf")
	if isDownload {
		w.Header().Set("Content-Disposition", "attachment; filename=laporan-operasional.pdf")
	} else {
		w.Header().Set("Content-Disposition", "inline; filename=laporan-operasional.pdf")
	}

	err = GeneratePDFLo(data, namaPT, bulan, w)
	if err != nil {
		http.Error(w, "Gagal membuat PDF: "+err.Error(), http.StatusInternalServerError)
	}
}

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
			"Rp. " + formatRupiah(row.Tarif),
			"Rp. " + formatRupiah(row.BiayaAngkut),
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
			pdf.CellFormat(widths[i], 5, "Rp. "+formatt(totalBiaya), "1", 0, "C", false, 0, "")
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
			pdf.CellFormat(widths[i], 5, "Rp. "+formatt(ppn), "1", 0, "R", false, 0, "")
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
			pdf.CellFormat(widths[i], 5, "Rp. "+formatt(grandTotal), "1", 0, "R", false, 0, "")
		default:
			pdf.CellFormat(widths[i], 5, " ", "1", 0, "", false, 0, "")
		}
	}
	pdf.Ln(30)
	pdf.SetFont("Arial", "", 8)
	pdf.Cell(0, 10, namaPT)

	return pdf.Output(w)
}

// Format angka seperti 12345678 menjadi 12.345.678
func formatt(num float64) string {
	rounded := int64(num)
	s := humanize.Comma(rounded)
	s = strings.Replace(s, ",", ".", -1)
	return s
}

func formatRupiah(num float64) string {
	s := humanize.CommafWithDigits(num, 2) // pakai 2 desimal
	s = strings.Replace(s, ",", "_", -1)
	s = strings.Replace(s, ".", ",", -1)
	s = strings.Replace(s, "_", ".", -1)
	return s
}
