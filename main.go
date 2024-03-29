package main

import (
	"fmt"
	"math/rand"
)

func main() {
	lowLetters := "abcdefghijklmnopqrstuvwxyz"
	highLetters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialChar := "!@3$%^&*()_+{}|[]/?><"
	numbers := "0123456789"

	var length int
	fmt.Println("Enter your Number:")
	_, err := fmt.Scan(&length)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	password := ""

	for n := 0; n < length; n++ {
		randomNum := rand.Intn(4)
		// fmt.Println(randomNum)

		switch randomNum {
		case 0:
			randomNum := rand.Intn(len(lowLetters))
			password = password + string(lowLetters[randomNum])
			break
		case 1:
			randomNum := rand.Intn(len(highLetters))
			password = password + string(highLetters[randomNum])
			break
		case 2:
			randomNum := rand.Intn(len(specialChar))
			password = password + string(specialChar[randomNum])
			break
		case 3:
			randomNum := rand.Intn(len(numbers))
			password = password + string(numbers[randomNum])
			break
		}
	}

	fmt.Println("\n\nYour password is:", password)

}
