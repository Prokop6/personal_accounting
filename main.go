package main

import (
	"fmt"

	_ "github.com/lib/pq"

	db_connections "github.com/Prokop6/personal-accounting/internal/db_connections"

	io_operations "github.com/Prokop6/personal-accounting/internal/io_operations"
)

func main() {
	db_connections.Init_db_connection()

	files, err := io_operations.List_files()

	if err != nil {
		panic(err)
	}

	for _, file := range files {

		transaction := io_operations.ReadYaml(file.Name(), io_operations.SourceFilesDir)

		isOk, expected, actual := transaction.Validate()

		if isOk {
			io_operations.Move_file(file.Name(), io_operations.SourceFilesDir, io_operations.ArchSubDir)
			fmt.Println(file.Name())
		} else {
			fmt.Printf("Found issue with file %s \n", file.Name())
			fmt.Printf("\tExpected value %f mismatches the actuall value of %f\n", expected, actual)
			io_operations.Move_file(file.Name(), io_operations.SourceFilesDir, io_operations.ErrorSubDir)

		}

	}
}

func vaildateFiles() {

}
