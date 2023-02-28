package helpers

import (
	"log"
	"os"
)

var Csv iCsv = csv{}

type iCsv interface {
	OpenFile() *os.File
}

type csv struct {
	iCsv
}

func (s csv) OpenFile() *os.File {
	file, err := os.Open("data/data.csv")
	if err != nil {
		log.Fatal(err)
	}

	return file
}
