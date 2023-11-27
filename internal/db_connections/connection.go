package db_connections

import (
	"database/sql"
	"fmt"
)

const (
	host     = "192.168.1.18"
	port     = 5432
	user     = "root"
	password = "12345"
	dbname   = "accounting"
)

func Init_db_connection() {

	connection_string := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", connection_string)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err)
	}

}
