package db_operations

import (
	"database/sql"

	"github.com/Prokop6/personal-accounting/internal/data_structures"
	utils "github.com/Prokop6/personal-accounting/internal/utils"
)

func CreateTransactionRecord(dbConnection *sql.DB, transaction *data_structures.Transaction, partnerID int32) int32 {


	transactionInsertStatement := `INSERT INTO public.transactions (date, partner_id, account, payment_method, currency, sum) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`

	queryResponse := dbConnection.QueryRow(transactionInsertStatement, transaction.GetDate(), partnerID, transaction.GetAccount(), transaction.GetMethod(), transaction.GetCurrency(), transaction.GetSum())

	var transactionID int32

	err := queryResponse.Scan(&transactionID)

			if err != nil {
				utils.Logger.Error(err)
				utils.Logger.Exit(1)
			}

	return transactionID	

}