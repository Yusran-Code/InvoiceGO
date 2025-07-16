package handlers

import (
	"invoice-go/service"
	"net/http"
)

func InvoicePDFHandler(w http.ResponseWriter, r *http.Request) {
	download := r.URL.Query().Get("download") == "true"
	err := service.GenerateInvoicePDF(w, r, download)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
