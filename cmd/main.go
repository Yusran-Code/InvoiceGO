package main

import (
	"dasar-go/auth"
	"dasar-go/config"
	"dasar-go/routes"

	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/dustin/go-humanize"
	"github.com/joho/godotenv"
)

func main() {
	// 1️⃣ Load .env terlebih dahulu
	if err := godotenv.Load(); err != nil {
		log.Println("Tidak bisa load .env:", err)
	} else {
		log.Println("✅ .env berhasil dimuat")
	}

	// 2️⃣ Debug ENV
	fmt.Println("DATABASE_URL:", os.Getenv("DATABASE_URL"))

	// 3️⃣ Template fungsi format
	tmpl := template.Must(template.New("").Funcs(template.FuncMap{
		"formatRupiah": func(n float64) string {
			return humanize.Comma(int64(n))
		},
	}).ParseGlob("templates/*.html"))

	// 4️⃣ Inisialisasi koneksi DB
	config.Init()
	defer config.DB.Close()

	// 5️⃣ Inisialisasi OAuth (setelah env ready)
	auth.InitOAuthConfig()

	// 6️⃣ Setup mux dan routing
	mux := http.NewServeMux()
	auth.RegisterAuthRoutes(mux)
	routes.RegisterAppRoutes(mux, tmpl, config.DB)

	log.Println("✅ Server berjalan di http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
