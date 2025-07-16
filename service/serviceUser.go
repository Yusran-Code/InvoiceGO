package service

import (
	"invoice-go/repository"
	"database/sql"
)

func IsUserProfileExist(db *sql.DB, email string) bool { // service
	_, err := repository.GetUserEmail(db, email)
	return err == nil
}
