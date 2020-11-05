package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "kspsql"
	password = "pass1111"
	dbname   = "csv_db"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, user, password, dbname)

	db, err := sql.Open("postgres", dataSourceName)
	checkError(err)

	defer db.Close()

	path, err := os.Getwd()
	checkError(err)

	files, err := ioutil.ReadDir(path + "/files")
	checkError(err)

	for _, file := range files {
		data, err := ioutil.ReadFile(path + "/files" + "/" + file.Name())
		checkError(err)

		reader := csv.NewReader(strings.NewReader(string(data)))

		_, err = reader.Read()
		checkError(err)

		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			}
			checkError(err)

			name := record[1]
			_, err = db.Exec("INSERT INTO names (name) VALUES ($1)", name)
			checkError(err)
		}
	}

	_, err = db.Exec(
		`DELETE FROM names x
		USING names y
		WHERE x.id < y.id
		AND x.name = y.name`)
	checkError(err)
}
