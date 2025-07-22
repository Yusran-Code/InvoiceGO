package service

import (
	"database/sql"
	"invoice-go/repository"
)

func IsUserProfileExist(db *sql.DB, email string) bool {
	_, err := repository.GetUserEmail(db, email)
	return err == nil
}
