package service

import (
	"invoice-go/model"
	"invoice-go/utils"
	"net/http"
	"strconv"
)

func ServiceInvoice(r *http.Request) (model.InvoiceData, error) {
	err := r.ParseForm()
	if err != nil {
		return model.InvoiceData{}, err
	}

	qty, _ := strconv.ParseFloat(r.FormValue("quantity_kg"), 64)
	dpp, _ := strconv.ParseFloat(r.FormValue("dpp"), 64)

	displayQty, pokok, ppn, dpp, total := utils.HitungTagihan(qty, dpp)

	data := model.InvoiceData{
		InvoiceNumber: r.FormValue("invoice_number"),
		InvoiceDate:   r.FormValue("invoice_date"),
		Periode:       r.FormValue("periode"),
		QuantityKG:    qty,
		DisplayQty:    displayQty,
		DPP:           dpp,
		Pokok:         pokok,
		PPN:           ppn,
		Total:         total,
	}

	return data, nil
}
