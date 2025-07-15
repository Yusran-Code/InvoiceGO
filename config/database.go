package config

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB // Ini untuk dipakai dari file lain: config.DB

func Init() {
	var err error
	connStr := os.Getenv("DATABASE_URL")

	maxRetries := 10
	for i := 0; i < maxRetries; i++ {
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			log.Println("Gagal buka koneksi DB:", err)
		} else {
			err = db.Ping()
			if err == nil {
				log.Println("Koneksi DB berhasil")
				DB = db // <-- SIMPAN ke global var
				return
			}
			log.Println("Ping DB gagal:", err)
			db.Close()
		}

		log.Printf("Coba konek DB lagi (%d/%d)...", i+1, maxRetries)
		time.Sleep(2 * time.Second)
	}

	log.Fatal("DB tidak merespon setelah beberapa percobaan:", err)
}
