package handlers

import (
	"invoice-go/service"
	"net/http"
	"time"
)

func HandleGeneratePDF(w http.ResponseWriter, r *http.Request, isDownload bool) {
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

	data, err := service.ParseExcelToDataRows(file)
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

	err = service.GeneratePDFLo(data, namaPT, bulan, w)
	if err != nil {
		http.Error(w, "Gagal membuat PDF: "+err.Error(), http.StatusInternalServerError)
	}
}
