package service

import (
	"dasar-go/model"
	"html/template"
	"net/http"
	"strconv"
)

func HandleGenerate(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			r.ParseForm()

			qty, _ := strconv.ParseFloat(r.FormValue("quantity_kg"), 64)
			dpp, _ := strconv.ParseFloat(r.FormValue("dpp"), 64)

			const hargaSatuan = 560.0
			const pengali = 3.0
			const hargaPokok = 595795.0

			displayQty := qty * hargaSatuan * pengali
			pokok := qty * hargaPokok
			ppn := pokok * 0.12
			total := pokok + ppn

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

			err := tmpl.ExecuteTemplate(w, "invoice.html", data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		tmpl.ExecuteTemplate(w, "generate.html", nil)
	}
}
