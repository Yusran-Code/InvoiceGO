package handlers

import (
	"invoice-go/service"
	"invoice-go/utils"
	"net/http"
)

func HandleGeneratePDF(w http.ResponseWriter, r *http.Request, isDownload bool) {
	form, err := utils.ParseInvoiceForm(r)
	if err != nil {
		http.Error(w, "Gagal parsing form: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer form.File.Close()

	data, err := service.ParseExcelToDataRows(form.File)
	if err != nil || len(data) == 0 {
		http.Error(w, "Gagal parsing data Excel", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	if isDownload {
		w.Header().Set("Content-Disposition", "attachment; filename=laporan-operasional.pdf")
	} else {
		w.Header().Set("Content-Disposition", "inline; filename=laporan-operasional.pdf")
	}

	err = service.GeneratePDFLo(data, form.NamaPT, form.Bulan, w)
	if err != nil {
		http.Error(w, "Gagal membuat PDF: "+err.Error(), http.StatusInternalServerError)
	}
}
