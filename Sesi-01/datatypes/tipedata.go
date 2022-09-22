package main

import "fmt"

func main(){
	// Data Types Pada Golang
	fmt.Println("===================================")
	fmt.Println("        Belajar Data Types         ")
	fmt.Println("===================================")

	// Tipe Data Number (integer)
	var first = 89
	var second = -12

	fmt.Println("         Tipe Data Integer (1)     ")
	fmt.Printf("\n")

	fmt.Printf("Tipe data first  : %T\n", first)
	fmt.Printf("Bilangan second  : %T\n", second)
	fmt.Println("===================================")
	fmt.Printf("\n")

	// Tipe Data Number  (integer)
	var numFirst uint8 = 89
	var numSecond int8 = -12

	fmt.Println("         Tipe Data Integer (2)     ")
	fmt.Printf("\n")

	fmt.Printf("Tipe data first  : %T\n", numFirst)
	fmt.Printf("Tipe data second : %T\n", numSecond)
	fmt.Println("===================================")
	fmt.Printf("\n")
	/*
		Kodingan tipe data integer 2 ini digunakan untuk
		menentukan tipe data integer dengan jenis apa yang
		ingin kita pakai untuk menghindari boros memori.
	*/

	// Tipe Data Number (float)
	var decimalNumber float32 = 3.52

	fmt.Println("           Tipe Data Float         ")
	fmt.Printf("\n")

	fmt.Printf("decimal number : %f\n", decimalNumber)
	fmt.Printf("decimal number : %.3f\n", decimalNumber)
	fmt.Println("===================================")
	fmt.Printf("\n")	
	/*
		verb %f digunakan untuk memformat angka desimal 
		atau tipe data float menjadi 6 digit angka desimal
		sedangkan verb %.nf digunakan untuk mengecilkan 
		digit desimalnya
	*/

	// Tipe Data Bool  
	var condition bool = true

	fmt.Println("           Tipe Data Bool         ")
	fmt.Printf("\n")

	fmt.Printf("Is it Permitted? %t \n", condition)
	fmt.Println("===================================")
	fmt.Printf("\n")	
	/*
		verb %t digunakan untuk memformat tipe data bool
		menjadi tipe data string
	*/


	// Tipe Data String
	var message string = `
	Selamat Datang di "Hacktiv8".
	Salam Kenal :).
	Mari belajar "Scalable Web Services With Go",
	`

	fmt.Println("           Tipe Data String        ")
	fmt.Printf("\n")

	fmt.Println(message)
	fmt.Println("===================================")
	fmt.Printf("\n")
}