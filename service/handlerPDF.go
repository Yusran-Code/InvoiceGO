package service

import (
	"dasar-go/auth"
	"dasar-go/config"
	"dasar-go/repository"
	"dasar-go/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dustin/go-humanize"
	"github.com/jung-kurt/gofpdf"
)

func HandleGeneratePDF(w http.ResponseWriter, r *http.Request) {
	session, _ := auth.GetSession(r)
	email := session.Values["email"].(string)

	profile, err := repository.LoadUserProfileByEmail(config.DB, email)
	if err != nil {
		http.Error(w, "Profil belum diisi", http.StatusInternalServerError)
		return
	}

	r.ParseForm()
	invoiceNumber := r.FormValue("invoice_number")
	invoiceDate := r.FormValue("invoice_date")
	qty, _ := strconv.ParseFloat(r.FormValue("quantity_kg"), 64)
	dpp, _ := strconv.ParseFloat(r.FormValue("dpp"), 64)
	periode := r.FormValue("periode")

	const hargaSatuan = 560.0
	const pengali = 3.0
	const hargaPokok = 595795.0

	displayQty := qty * hargaSatuan * pengali
	pokok := qty * hargaPokok
	ppn := pokok * 0.12
	total := pokok + ppn

	// === PDF generation ===
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(20, 20, 20)
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 14)
	pdf.CellFormat(0, 7, profile.NamaPT, "", 1, "C", false, 0, "")
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(0, 5, "Agen LPG PSO", "", 1, "C", false, 0, "")
	pdf.CellFormat(0, 5, profile.Alamat+" , "+profile.Kabupaten, "", 1, "C", false, 0, "")
	pdf.Ln(5)

	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(0, 10, "INVOICE", "", 1, "C", false, 0, "")
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 9)
	pdf.Cell(0, 6, "Kepada : PT.Pertamina Patra Niaga")
	pdf.Ln(6)
	pdf.Cell(0, 6, "Alamat : Gedung Wisma Tugu II Lt.2, Jl. HR Rasuna Said KAV C7-9 Setiabudi, Jakarta 12920")
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(95, 6, "Tanggal: "+invoiceDate, "", 0, "L", false, 0, "")
	pdf.SetX(115)
	pdf.CellFormat(0, 6, "No. Invoice: "+invoiceNumber, "", 1, "R", false, 0, "")
	pdf.Ln(5)

	pdf.SetFont("Arial", "B", 11)
	pdf.SetFillColor(240, 240, 240)
	pdf.CellFormat(90, 8, "Deskripsi", "1", 0, "C", true, 0, "")
	pdf.CellFormat(90, 8, "Nilai", "1", 1, "C", true, 0, "")

	pdf.SetFont("Arial", "", 11)
	pdf.SetFillColor(255, 255, 255)
	pdf.CellFormat(90, 8, "Tagihan Transport Fee Periode "+periode, "1", 0, "L", false, 0, "")
	pdf.CellFormat(90, 8, "", "1", 1, "R", false, 0, "")
	pdf.CellFormat(90, 8, "Quantity/Kg", "1", 0, "L", false, 0, "")
	pdf.CellFormat(90, 8, humanize.Comma(int64(displayQty)), "1", 1, "R", false, 0, "")
	pdf.CellFormat(90, 8, "DPP", "1", 0, "L", false, 0, "")
	pdf.CellFormat(90, 8, humanize.Comma(int64(dpp)), "1", 1, "R", false, 0, "")
	pdf.CellFormat(90, 8, "Pokok", "1", 0, "L", false, 0, "")
	pdf.CellFormat(90, 8, humanize.Comma(int64(pokok)), "1", 1, "R", false, 0, "")
	pdf.CellFormat(90, 8, "PPN 12%", "1", 0, "L", false, 0, "")
	pdf.CellFormat(90, 8, humanize.Comma(int64(ppn)), "1", 1, "R", false, 0, "")

	pdf.SetFont("Arial", "B", 11)
	pdf.CellFormat(90, 8, "Total", "1", 0, "L", false, 0, "")
	pdf.CellFormat(90, 8, "Rp. "+humanize.Comma(int64(total)), "1", 1, "R", false, 0, "")

	pdf.Ln(8)
	pdf.SetFont("Arial", "I", 10)
	pdf.MultiCell(0, 6, "Terbilang: "+utils.Terbilang(int64(total)), "", "", false)
	fmt.Println(utils.Terbilang(123456))

	pdf.SetFont("Arial", "", 10)
	pdf.Ln(15)
	pdf.Cell(0, 6, "Bank: "+profile.NamaBank+" - "+profile.NoRekening)
	pdf.Ln(6)
	pdf.Cell(0, 6, "a.n "+profile.NamaPT)

	pdf.Ln(35)
	pdf.CellFormat(0, 6, profile.PenanggungJawab, "", 1, "R", false, 0, "")

	download := r.URL.Query().Get("download") == "true"
	w.Header().Set("Content-Type", "application/pdf")
	if download {
		w.Header().Set("Content-Disposition", "attachment; filename=invoice.pdf")
	} else {
		w.Header().Set("Content-Disposition", "inline; filename=invoice.pdf")
	}
	if err := pdf.Output(w); err != nil {
		http.Error(w, "Gagal membuat PDF", http.StatusInternalServerError)
	}
}
