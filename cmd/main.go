package main

import (
	"invoice-go/auth"
	"invoice-go/config"
	"invoice-go/routes"

	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/dustin/go-humanize"
	"github.com/joho/godotenv"
)

func main() {
	// Coba load .env, tapi tidak panik jika gagal (bisa karena Railway)
	if err := godotenv.Load(".env"); err != nil {
		log.Println("🔔 Tidak bisa load .env (mungkin karena running di Railway):", err)
	} else {
		log.Println("✅ .env berhasil dimuat")
	}

	// Pastikan SESSION_KEY tetap ada meskipun dari environment Railway
	if os.Getenv("SESSION_KEY") == "" {
		log.Fatal("❌ SESSION_KEY belum di-set! Set di .env atau di Railway environment variables.")
	}

	// Inisialisasi session
	auth.InitSession()

	// Tambahkan fungsi format Rupiah ke template
	tmpl := template.Must(template.New("").Funcs(template.FuncMap{
		"formatRupiah": func(n float64) string {
			return humanize.Comma(int64(n))
		},
	}).ParseGlob("templates/*.html"))

	// Koneksi ke DB
	config.Init()
	defer config.DB.Close()

	// Inisialisasi OAuth setelah env
	auth.InitOAuthConfig()

	// Routing
	mux := http.NewServeMux()
	auth.RegisterAuthRoutes(mux)
	routes.RegisterAppRoutes(mux, tmpl, config.DB)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("✅ Server berjalan di http://localhost:%s\n", port)

	err := http.ListenAndServe("0.0.0.0:"+port, mux)
	if err != nil {
		log.Fatalf("❌ Gagal menjalankan server: %v", err)
	}

}
