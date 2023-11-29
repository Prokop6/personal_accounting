package db_operations

import (
	"database/sql"

	"github.com/Prokop6/personal-accounting/internal/data_structures"
	"github.com/Prokop6/personal-accounting/internal/utils"
)


func CreateTransactionItemRecord (dbConn *sql.DB, transID int32, transItem data_structures.Items) {
	sqlStatement := `INSERT INTO public.transaction_items (transaction_id, name, ammount, unit_price) VALUES ($1, $2, $3, $4)`

	_, err := dbConn.Exec(sqlStatement,transID, transItem.Name, transItem.Amount, transItem.Price)  

	if err != nil {
		utils.Logger.Error(err)
		utils.Logger.Exit(1)
	}

}