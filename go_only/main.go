package main

import (
	"encoding/csv"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func foundElement(name string, result []string) bool {
	for _, value := range result {
		if value == name {
			return true
		}
	}
	return false
}

func main() {
	path, err := os.Getwd()
	checkError(err)

	files, err := ioutil.ReadDir(path + "/files")
	checkError(err)

	result := make([]string, 0)

	for _, file := range files {

		data, err := ioutil.ReadFile(path + "/files/" + file.Name())
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
			isExitElement := foundElement(name, result)

			if isExitElement == false {
				result = append(result, name)
			}
		}

	}

	f, err := os.OpenFile(path+"/result.csv", os.O_APPEND|os.O_WRONLY, 0644)
	checkError(err)
	defer f.Close()

	for _, value := range result {
		_, err = f.WriteString(value + "\n")
		checkError(err)
	}

}
