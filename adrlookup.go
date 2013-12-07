package main

import (
	"fmt"
	"encoding/csv"
	"os"
)

func main () {
	loadADRCodes("adrUNno.csv")
}

func loadADRCodes(filename string)  {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	r := csv.NewReader(file)
	r.Comma = '|'
	r.FieldsPerRecord = -1

	row, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(row); i++ {
		fmt.Println(row[i][1])
	}
}
