package io_operations

import (
	"io/fs"
	"os"
	"path/filepath"
)

const SourceFilesDir = "/workspaces/personal_accounting/inputs"
const ArchSubDir = "arch"
const ErrorSubDir = "errors"

func Move_file(fileName string, sourceDir string, targetSubDir string) {
	from := filepath.Join(sourceDir, fileName)
	to := filepath.Join(sourceDir, targetSubDir, fileName)

	os.Rename(from, to)
}


func List_files() ([]fs.DirEntry, error) {
	dir, err := os.Open(SourceFilesDir)



	if err != nil {

		return nil, err
	}

	files, err := dir.ReadDir(0)

	if err != nil {

		return nil, err
	}

	file_list := make([]fs.DirEntry, 0)

	for _, file := range files {

		if file.IsDir() {

			continue
		}

		file_list = append(file_list, file)

	}

	return file_list, nil

}