package models

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type Zip struct {
	// uppercase auto exported
	Code  string
	City  string
	State string
}

type ZipSlice []*Zip // pointer to the Zip struct

type ZipIndex map[string]ZipSlice

func LoadZips(fileName string) (ZipSlice, error) { // can have multiple return types
	// Read in a csv file and return to the caller
	f, err := os.Open(fileName) // file path
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}

	reader := csv.NewReader(f) // will read one line at a time and process as string
	_, err = reader.Read()     // ignore column headings
	if err != nil {
		return nil, fmt.Errorf("error reading header row: %v", err)
	}

	// pre-allocating the under lying array to be 43000 long
	zips := make(ZipSlice, 0, 43000)
	for {
		// returns a slice of strings or an error if there was a problem
		fields, err := reader.Read()
		if err == io.EOF { // when it is done iterating
			return zips, nil
		}
		if err != nil {
			return nil, fmt.Errorf("error reading record: %v", err)
		}
		z := &Zip{
			Code:  fields[0],
			City:  fields[3],
			State: fields[6],
		}
		zips = append(zips, z) // append into a slice and return new slice
	}
}
