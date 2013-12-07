package main

import (
	"fmt"
	"encoding/csv"
	"os"
)

func loadADRCodes(filename string)  {
	file, err := os.Open(filename)
    	if err != nil {
		panic(err)
	}
	defer file.Close()
	r := csv.NewReader(file)

	row, e := r.Read()
	if e != nil {
		panic(e)
	}
	fmt.Println(row)
}
