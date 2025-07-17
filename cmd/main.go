package main

import (
	"invoice-go/auth"
	"invoice-go/config"
	"invoice-go/routes"

	"html/template"
	"log"
	"net/http"

	"github.com/dustin/go-humanize"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env terlebih dahulu
	if err := godotenv.Load(); err != nil {
		log.Println("Tidak bisa load .env:", err)
	} else {
		log.Println("✅ .env berhasil dimuat")
	}
	auth.InitSession()

	//  Template fungsi format
	tmpl := template.Must(template.New("").Funcs(template.FuncMap{
		"formatRupiah": func(n float64) string {
			return humanize.Comma(int64(n))
		},
	}).ParseGlob("templates/*.html"))

	//  Inisialisasi koneksi DB
	config.Init()
	defer config.DB.Close()

	//  Inisialisasi OAuth (setelah env ready)
	auth.InitOAuthConfig()

	//  Setup mux dan routing
	mux := http.NewServeMux()
	auth.RegisterAuthRoutes(mux)
	routes.RegisterAppRoutes(mux, tmpl, config.DB)

	log.Println("✅ Server berjalan di http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
