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
	tmpl := template.Must(template.New("").Funcs(template.FuncMap{
		"formatRupiah": func(n float64) string {
			return humanize.Comma(int64(n))
		},
	}).ParseGlob("templates/*.html"))

	if os.Getenv("ENV") != "production" {
		env := os.Getenv("ENV")
		fmt.Printf("✅ Connected to PostgreSQL (%s environment)\n", env)

		_ = godotenv.Load()
	}

	//init
	//config.LoadEnv()
	auth.InitOAuthConfig()
	config.Init()
	defer config.DB.Close()

	mux := http.NewServeMux()

	// ✅ Routing Auth dan App
	auth.RegisterAuthRoutes(mux)
	routes.RegisterAppRoutes(mux, tmpl, config.DB)

	//hadnler LO

	// ✅ Static Handler for /static/

	log.Println("Jalan di http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
