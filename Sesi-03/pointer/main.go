package main

import (
	"fmt"
	"strings"
)

func main() {
	// Deklarasi Pointer
	fmt.Println("====================================")
	fmt.Println("           Memory Address           ")
	fmt.Println("====================================")

	var firstNumber int = 4
	var secondNumber *int = &firstNumber

	fmt.Println("firstNumber (value)         :", firstNumber)
	fmt.Println("firstNumber (memori addres) :", &firstNumber)

	fmt.Println("secondNumber (value)        :", *secondNumber)
	fmt.Println("secondNumber (memori addres):", &secondNumber)

	fmt.Println()

	// Mengubah Nilai Pada Pointer
	fmt.Println("====================================")
	fmt.Println("   Changing Value through Pointer   ")
	fmt.Println("====================================")

	var firstPerson string = "John"
	var secondPerson *string = &firstPerson

	fmt.Println("firstPerson (value)          :", firstPerson)
	fmt.Println("firstPerson (memory address) :", &firstPerson)
	fmt.Println("secondPerson (value)         :", *secondPerson)
	fmt.Println("secondPerson (memory address):", &secondPerson)

	fmt.Println("\n", strings.Repeat("#", 30), "\n")

	*secondPerson = "Doe"

	fmt.Println("firstPerson (value)          :", firstPerson)
	fmt.Println("firstPerson (memory address) :", &firstPerson)
	fmt.Println("secondPerson (value)         :", *secondPerson)
	fmt.Println("secondPerson (memory address):", &secondPerson)

	fmt.Println()

	// Parameter Pointer
	fmt.Println("====================================")
	fmt.Println("          Parameter Pointer         ")
	fmt.Println("====================================")

	var number = 4
	fmt.Println("before :", number) // 4
	change(&number, 10)
	fmt.Println("after  :", number) // 10

	fmt.Println()
}

func change(original *int, value int) {
	*original = value
}
