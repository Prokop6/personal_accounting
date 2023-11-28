package db_connections

import (
	"database/sql"
	"fmt"

	"github.com/Prokop6/personal-accounting/internal/utils"
)

const (
	host     = "192.168.1.18"
	port     = 5432
	user     = "root"
	password = "12345"
	dbname   = "accounting"
)

func InitDBCconnection() (*sql.DB) {

	connection_string := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", connection_string)

	if err != nil {
		utils.Logger.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		utils.Logger.Fatal(err)
	}

	return db

}
