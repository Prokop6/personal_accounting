package main

import (
	"fmt"

	_ "github.com/lib/pq"

	db_connections "github.com/Prokop6/personal-accounting/internal/db_connections"

	io_operations "github.com/Prokop6/personal-accounting/internal/io_operations"

	utils "github.com/Prokop6/personal-accounting/internal/utils"
)

func main() {

	utils.InitLogger()

	dbConnection := db_connections.InitDBCconnection()

	files, err := io_operations.List_files()

	if err != nil {
		panic(err)
	}

	sqlStatement := `INSERT INTO public.transactions (date, partner, account, payment_method, currency, sum) VALUES ($1, $2, $3, $4, $5, $6);`

	for _, file := range files {

		transaction := io_operations.ReadYaml(file.Name(), io_operations.SourceFilesDir)

		isOk, expected, actual := transaction.Validate()

		if isOk {
			utils.Logger.Info("Adding record to database")

			_, err = dbConnection.Exec(sqlStatement, transaction.GetDate(), transaction.GetPartnerName(), transaction.GetAccount(), transaction.GetMethod(), transaction.GetCurrency(), transaction.GetSum())

			if err != nil {
				utils.Logger.Error(err)
				utils.Logger.Exit(1)
			}

			io_operations.Move_file(file.Name(), io_operations.SourceFilesDir, io_operations.ArchSubDir)
			fmt.Println(file.Name())

		} else {
			utils.Logger.Errorf("Found issue with file %s \n", file.Name())
			utils.Logger.Errorf("\tExpected value %f mismatches the actuall value of %f\n", expected, actual)
			io_operations.Move_file(file.Name(), io_operations.SourceFilesDir, io_operations.ErrorSubDir)

		}

		defer dbConnection.Close()


	}
}

func vaildateFiles() {

}
