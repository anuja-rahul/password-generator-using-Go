package main

import (
	"encoding/csv"
	"log"
	"os"
)

func checkError(err error) {
	if err != nil {
		log.Panicln(err)
	}
}

const FILENAME string = "passwords.csv"

func readFile() ([][]string, error) {

	f, err := os.Open(FILENAME)
	if err != nil {
		f, err = os.Create(FILENAME)
		checkError(err)
		log.Println("Created FIle !")
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			checkError(err)
		}
	}(f)

	r := csv.NewReader(f)
	entries, err := r.ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return entries, nil
}

func main() {

	// fmt.Printf("file: %v\n", f)
}
