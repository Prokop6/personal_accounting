package io_operations

import (
	"os"
	"path/filepath"

	"github.com/Prokop6/personal-accounting/internal/data_structures"
	"gopkg.in/yaml.v3"
)

func ReadYaml(fileName string, dirName string) *data_structures.Transaction {

	var data data_structures.Transaction

	fileFullPath := filepath.Join(dirName, fileName)
	file, err := os.ReadFile(fileFullPath)

	if err != nil {
		Move_file(fileName, dirName, ErrorSubDir)
		panic(err)
	}

	err = yaml.Unmarshal(file, &data)

	if err != nil {
		Move_file(fileName, dirName, ErrorSubDir)

		panic(err)

	}

	return &data

}
