package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	var number int
	var err error

	number, err = strconv.Atoi("123GH")

	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	number, err = strconv.Atoi("123")

	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	// Memanggil func catchErr()
	defer catchErr()



	// Custom Error
	var password string
	
	fmt.Scanln(&password)

	if valid, err := validPassword(password); err != nil {
		// Panic
		panic(err.Error())
	} else {
		fmt.Println(valid)
	}
}

// Recover
func catchErr() {
	if r := recover(); r != nil {
		fmt.Println("Error Ocured:", r)
	} else {
		fmt.Println("Application Running Perfectly")
	}
}

func validPassword(password string) (string, error) {
	pl := len(password)

	if pl < 5 {
		return "", errors.New("Password Has to Have more than 4 characters")
	}

	return "Valid Password", nil
}