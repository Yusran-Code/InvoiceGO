package utils

import (
	"mime/multipart"
	"net/http"
	"time"
)

type InvoiceForm struct {
	NamaPT string
	Bulan  string
	File   multipart.File
}

func ParseInvoiceForm(r *http.Request) (InvoiceForm, error) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		return InvoiceForm{}, err
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		return InvoiceForm{}, err
	}

	namaPT := r.FormValue("namapt")
	bulan := r.FormValue("bulan")
	if bulan == "" {
		bulan = time.Now().Format("January 2006")
	}

	return InvoiceForm{
		NamaPT: namaPT,
		Bulan:  bulan,
		File:   file,
	}, nil
}
