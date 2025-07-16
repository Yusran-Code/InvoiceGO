package service

import (
	"invoice-go/model"
	"invoice-go/repository"
	"database/sql"
)

func LoadProfileByEmail(db *sql.DB, email string) (*model.AppProfile, error) {
	return repository.GetUserEmail(db, email)
}

func UpdateProfile(db *sql.DB, profile model.AppProfile) error {
	return repository.SaveUserProfile(db, profile)
}
