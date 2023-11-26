package main

import (
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"

	data_structures "github.com/Prokop6/personal-accounting/internal/data_structures"

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

		_ = readYaml(file.Name(), io_operations.SourceFilesDir)

		io_operations.Move_file(file.Name(), io_operations.SourceFilesDir, io_operations.ArchSubDir)
		fmt.Println(file.Name())
	}
}

func vaildateFiles() {

}

func readYaml(fileName string, dirName string) *data_structures.Transaction {

	var data data_structures.Transaction

	fileFullPath := filepath.Join(dirName, fileName)
	file, err := os.Open(fileFullPath)

	if err != nil {
		io_operations.Move_file(fileName, dirName, io_operations.ErrorSubDir)
		panic(err)
	}

	var fileContent []byte

	_, err = file.Read(fileContent)

	if err != nil {
		io_operations.Move_file(fileName, dirName, io_operations.ErrorSubDir)
		panic(err)
	}

	err = yaml.Unmarshal(fileContent, &data)

	if err != nil {
		io_operations.Move_file(fileName, dirName, io_operations.ErrorSubDir)

		panic(err)

	}

	return &data

}
