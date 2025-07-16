package repository

import (
	"invoice-go/model"
	"database/sql"
)

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
