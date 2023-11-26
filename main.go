package main

import (
	"database/sql"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
)

const (
	host     = "192.168.1.18"
	port     = 5432
	user     = "root"
	password = "12345"
	dbname   = "accounting"
)

const sourceFilesDir = "/workspaces/personal_accounting/inputs"
const archSubDir = "arch"
const errorSubDir = "errors"

func init_db_connection() {

	connection_string := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", connection_string)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err)
	}

}

func list_files() ([]fs.DirEntry, error) {
	dir, err := os.Open(sourceFilesDir)

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

func main() {
	init_db_connection()

	files, err := list_files()

	if err != nil {
		panic(err)
	}

	for _, file := range files {

		_ = readYaml(file.Name(), sourceFilesDir)


		move_file(file.Name(), sourceFilesDir, archSubDir)
		fmt.Println(file.Name())
	}
}

func move_file(fileName string, sourceDir string, targetSubDir string) {
	from := filepath.Join(sourceDir, fileName)
	to := filepath.Join(sourceDir, targetSubDir, fileName)

	os.Rename(from, to)

}

func vaildateFiles() {

}

func readYaml(fileName string, dirName string) *Transaction {

	var data Transaction
	
	fileFullPath := filepath.Join(dirName, fileName)
	file, err :=	os.Open(fileFullPath)

	if err != nil {
		move_file(fileName, dirName, errorSubDir)
		panic(err)
	}

	var fileContent	[]byte

	_, err = file.Read(fileContent)

	if err != nil {
		move_file(fileName, dirName, errorSubDir)
		panic(err)
	}

	err = yaml.Unmarshal(fileContent, &data)

	if err != nil {
		move_file(fileName, dirName, errorSubDir)

		panic(err)

	}

	return &data

}

type Transaction struct {
	Date string `yaml:"date"`
	Shop string `yaml:"shop"`
	Account string `yaml:"account"`
	Method string `yaml:"method"`
	Sum string `yaml:"sum"`
	Items []Items `yaml:"items"`
}

type Items struct {
	Name string `yaml:"name"`
	Amount string `yaml:"amount"`
	Price string `yaml:"price"`
}