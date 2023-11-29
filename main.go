package main

import (
	_ "github.com/lib/pq"

	db_connections "github.com/Prokop6/personal-accounting/internal/db_connections"
	"github.com/Prokop6/personal-accounting/internal/db_operations"

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

	for _, file := range files {

		transaction := io_operations.ReadYaml(file.Name(), io_operations.SourceFilesDir)

		isOk, expected, actual := transaction.Validate()

		if isOk {
			utils.Logger.Info("Adding record to database")

			partnerID := db_operations.GetPartnerID(dbConnection, transaction.GetPartnerName())

			utils.Logger.Info("Creating record for transaction")

			transactionID := db_operations.CreateTransactionRecord(dbConnection, transaction, partnerID)

			utils.Logger.Infof("%d", transactionID)

			utils.Logger.Infof("Moving file %s to archive", file.Name())
			io_operations.Move_file(file.Name(), io_operations.SourceFilesDir, io_operations.ArchSubDir)

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
