package services

import "database/sql"

type Services struct {
	db *sql.DB
}

func InitServices(db *sql.DB) Services {
	return Services{
		db: db,
	}
}
