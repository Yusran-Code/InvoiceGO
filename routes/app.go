package routes

import (
	"dasar-go/auth"
	"dasar-go/service"
	"database/sql"
	"html/template"
	"net/http"
)

// daftar semua route
func RegisterAppRoutes(mux *http.ServeMux, tmpl *template.Template, db *sql.DB) {
	mux.HandleFunc("/", service.HandleHome(tmpl))
	mux.HandleFunc("/index", auth.RequireAuth(service.HandleIndex(tmpl, db)))
	mux.HandleFunc("/setup", auth.RequireAuth(service.HandleSetup(tmpl)))
	mux.HandleFunc("/generate", auth.RequireAuth(service.HandleGenerate(tmpl)))
	mux.HandleFunc("/generate-pdf", auth.RequireAuth(service.HandleGeneratePDF))
}
