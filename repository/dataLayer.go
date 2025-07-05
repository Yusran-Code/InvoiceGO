package repository

import (
	"dasar-go/model"
	"database/sql"
)

// Ambil profil berdasarkan email
func LoadUserProfileByEmail(db *sql.DB, email string) (model.AppProfile, error) {
	var p model.AppProfile
	err := db.QueryRow(`
		SELECT email, nama_pt, nama_bank, no_rekening, penanggung_jawab, alamat, kabupaten
		FROM user_profile WHERE email = $1
	`, email).Scan(&p.Email, &p.NamaPT, &p.NamaBank, &p.NoRekening, &p.PenanggungJawab, &p.Alamat, &p.Kabupaten)
	return p, err
}

// Simpan atau update profil berdasarkan email (UPSERT)
func SaveUserProfile(db *sql.DB, p model.AppProfile) error {
	_, err := db.Exec(`
		INSERT INTO user_profile (email, nama_pt, nama_bank, no_rekening, penanggung_jawab, alamat, kabupaten)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (email) DO UPDATE SET
			nama_pt = EXCLUDED.nama_pt,
			nama_bank = EXCLUDED.nama_bank,
			no_rekening = EXCLUDED.no_rekening,
			penanggung_jawab = EXCLUDED.penanggung_jawab,
			alamat = EXCLUDED.alamat,
			kabupaten = EXCLUDED.kabupaten
	`, p.Email, p.NamaPT, p.NamaBank, p.NoRekening, p.PenanggungJawab, p.Alamat, p.Kabupaten)
	return err
}
