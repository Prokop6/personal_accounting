package db_operations

import (
	"database/sql"

	utils "github.com/Prokop6/personal-accounting/internal/utils"
)

func GetPartnerID(con *sql.DB, name string) int32 {

	partnerIDStatement := `SELECT find_or_create_partner($1)`

	var partner_id int32

	queryResult := con.QueryRow(partnerIDStatement, name)

	err := queryResult.Scan(&partner_id) 

	if err != nil {
		utils.Logger.Error(err)
		utils.Logger.Exit(1)
	}

	return partner_id
}
