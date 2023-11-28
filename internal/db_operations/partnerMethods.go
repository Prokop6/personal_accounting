package db_operations

import (
	"database/sql"

	utils "github.com/Prokop6/personal-accounting/internal/utils"
)

func GetPartnerID(con *sql.DB, name string) int32 {

	partnerSelectStatement := `SELECT id, short_name FROM accounting.public.partners WHERE short_name = $1`

	partnerInsertStatement := `INSERT INTO public.partners
	(short_name) SELECT $1 WHERE NOT EXISTS (SELECT short_name FROM public.partners WHERE short_name = $1) RETURNING id, short_name`

	utils.Logger.Info("Searching for partner ID")

	queryRes := con.QueryRow(partnerSelectStatement,
		name)

	var partnerID int32
	var partnerShort string
	err := queryRes.Scan(&partnerID, &partnerShort)

	if err == sql.ErrNoRows {
		statementResult, err := con.Query(partnerInsertStatement, name)

		if err != nil {
			utils.Logger.Error(err)
			utils.Logger.Exit(1)
		}

		statementResult.Next()
		err = statementResult.Scan(&partnerID, &partnerShort)

		if err != nil {
			utils.Logger.Error(err)
			utils.Logger.Exit(1)
		}

		utils.Logger.Infof("Created record for %s with id %d", partnerShort,
			partnerID)

	} else {

		utils.Logger.Infof("Found record for %s with id %d", partnerShort,
			partnerID)

	}

	return partnerID
}
