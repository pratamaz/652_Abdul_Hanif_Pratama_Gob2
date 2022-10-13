package main

import "fmt"

func main() {
	// Belajar Constant
	fmt.Println("===================================")
	fmt.Println("        Belajar Constant           ")
	fmt.Println("===================================")

	// Constant
	const full_name string = "Abdul Hanif Pratama"

	fmt.Printf("Hello, %s\n", full_name)
	fmt.Println("===================================")
	fmt.Printf("\n")

	/*
		Ketika membuat variabel dengan const, maka harus
		langsung memberikan nilai pada variabel tersebut.
	*/

	// Belajar Operators
	fmt.Println("===================================")
	fmt.Println("         Belajar Operators         ")
	fmt.Println("===================================")

	// Operator Aritmetika
	var value1 = 2 + 2*3
	var value2 = (2 + 2) * 3

	fmt.Println("        Operator Aritmetika       ")
	fmt.Printf("\n")

	fmt.Println("Hasil :", value1)
	fmt.Println("Hasil :", value2)
	fmt.Println("===================================")
	fmt.Printf("\n")
	
	/*
		Pada nilai value1, Go akan melakukan perkalian
		terlebih dahulu kemudian baru melakukan pertambahan.
		Namun, jika kita ingin kalkulasi pertambahannya
		dieksekusi terlebih dahulu, maka kita bisa 
		mengelompokkannya dengan menggunakan simbol ().
	*/


	// Operator Relasional
	var firstCondition bool  = 2 < 3
	var secondCondition bool = "joey" == "Joey"
	var thirdCondition bool  = 10 != 2.3
	var fourthCondition bool = 11 <= 11

	fmt.Println("        Operator Relasional       ")
	fmt.Printf("\n")

	fmt.Println("First Condition  :", firstCondition)
	fmt.Println("Second Condition :", secondCondition)
	fmt.Println("Third Condition  :", thirdCondition)
	fmt.Println("Fourth Condition :", fourthCondition)
	
	fmt.Println("===================================")
	fmt.Printf("\n")


	// Operator Logika
	var right = true
	var wrong  = false 

	fmt.Println("          Operator Logika         ")
	fmt.Printf("\n")

	var wrongAndRight = wrong && right
	fmt.Printf("Wrong && right \t(%t) \n", wrongAndRight)
	
	var wrongOrRight = wrong || right
	fmt.Printf("Wrong || right \t(%t) \n", wrongOrRight)

	var wrongReverse = !wrong
	fmt.Printf("!Wrong \t\t(%t) \n", wrongReverse)
	
	fmt.Println("===================================")
	fmt.Printf("\n")
}
