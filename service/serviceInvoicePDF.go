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

	profile, err := repository.GetUserEmail(config.DB, email)
	if err != nil {
		return fmt.Errorf("profil belum diisi")
	}

	err = r.ParseForm()
	if err != nil {
		return fmt.Errorf("gagal parsing form: %v", err)
	}

	invoiceNumber := r.FormValue("invoice_number")
	invoiceDate := r.FormValue("invoice_date")
	periode := r.FormValue("periode")
	qty, _ := strconv.ParseFloat(r.FormValue("quantity_kg"), 64)
	dpp, _ := strconv.ParseFloat(r.FormValue("dpp"), 64)

	// 💰 Hitung tagihan manual via utils
	displayQty, pokok, ppn, _, total := utils.HitungTagihan(qty, dpp)

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

	pdf := utils.GeneratePDFInvoice(*profile, data)

	w.Header().Set("Content-Type", "application/pdf")
	if isDownload {
		w.Header().Set("Content-Disposition", "attachment; filename=invoice.pdf")
	} else {
		w.Header().Set("Content-Disposition", "inline; filename=invoice.pdf")
	}

	return pdf.Output(w)
}
