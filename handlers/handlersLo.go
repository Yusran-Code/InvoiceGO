package handlers

import (
	"html/template"
	"net/http"
)

// handlers/lo.go
func ShowLoPage(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "lo.html", nil)
	}
}

func PreviewLoHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "gunakan metode POST", http.StatusMethodNotAllowed)
			return
		}
		HandleGeneratePDF(w, r, false)
	}
}

func DownloadLoHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "gunakan metode POST", http.StatusMethodNotAllowed)
			return
		}
		HandleGeneratePDF(w, r, true)
	}
}
