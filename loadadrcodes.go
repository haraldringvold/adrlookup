package main

import (
	"encoding/csv"
	"os"
)

func loadADRCodes(filename string) []ADRCode {
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
	codes := make([]ADRCode, len(row))
	for i, v := range row {
		codes[i] = ADRCode{v[1], v[2], v[3], v[4], v[5], v[6], v[7], ""}
	}
	return codes
}
