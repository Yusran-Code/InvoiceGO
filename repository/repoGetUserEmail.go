package repository

import (
	"invoice-go/model"
	"database/sql"
)

func GetUserEmail(db *sql.DB, email string) (*model.AppProfile, error) {
	row := db.QueryRow(`
		SELECT email, nama_pt, nama_bank, no_rekening, penanggung_jawab, alamat, kabupaten
		FROM user_profile
		WHERE email = $1
	`, email)

	var profile model.AppProfile
	err := row.Scan(
		&profile.Email,
		&profile.NamaPT,
		&profile.NamaBank,
		&profile.NoRekening,
		&profile.PenanggungJawab,
		&profile.Alamat,
		&profile.Kabupaten,
	)
	if err != nil {
		return nil, err
	}

	return &profile, nil
}
