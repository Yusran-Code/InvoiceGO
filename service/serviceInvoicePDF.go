package service

import (
	"fmt"

	"invoice-go/auth"
	"invoice-go/config"
	"invoice-go/model"
	"invoice-go/repository"
	"invoice-go/utils"
	"net/http"
	"strconv"
)

func GenerateInvoicePDF(w http.ResponseWriter, r *http.Request, isDownload bool) error {
	session, _ := auth.GetSession(r)
	email := session.Values["email"].(string)

	// 🔍 Ambil profil user dari DB
	profile, err := repository.GetUserEmail(config.DB, email)
	if err != nil {
		return fmt.Errorf("profil belum diisi")

	}

	// 📝 Ambil input form
	err = r.ParseForm()
	if err != nil {
		return fmt.Errorf("gagal parsing form: %v", err)
	}

	invoiceNumber := r.FormValue("invoice_number")
	invoiceDate := r.FormValue("invoice_date")
	periode := r.FormValue("periode")
	qty, _ := strconv.ParseFloat(r.FormValue("quantity_kg"), 64)
	dpp, _ := strconv.ParseFloat(r.FormValue("dpp"), 64)

	// 💰 Hitung tagihan
	displayQty, pokok, ppn, total := utils.HitungTagihan(qty)

	// 📦 Bungkus ke struct model.InvoiceData
	data := model.InvoiceData{
		InvoiceNumber: invoiceNumber,
		InvoiceDate:   invoiceDate,
		Periode:       periode,
		QuantityKG:    qty,
		DisplayQty:    displayQty,
		DPP:           dpp,
		Pokok:         pokok,
		PPN:           ppn,
		Total:         total,
	}

	// 📄 Generate PDF
	pdf := utils.GeneratePDFInvoice(*profile, data)

	// 📤 Set header PDF
	w.Header().Set("Content-Type", "application/pdf")
	if isDownload {
		w.Header().Set("Content-Disposition", "attachment; filename=invoice.pdf")
	} else {
		w.Header().Set("Content-Disposition", "inline; filename=invoice.pdf")
	}

	// ✅ Kirim PDF ke browser
	return pdf.Output(w)
}
