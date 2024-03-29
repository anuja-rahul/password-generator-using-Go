package main

import (
	"fmt"
	"log"
	"os"
)

func checkError(err error) {
	if err != nil {
		log.Panicln(err)
	}
}

const FILENAME string = "passwords.csv"

func main() {
	f, err := os.Open(FILENAME)
	if err != nil {
		f, err = os.Create(FILENAME)
		checkError(err)
		fmt.Println("Created FIle !")
	}
	fmt.Printf("file: %v\n", f)
}
