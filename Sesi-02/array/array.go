package main

import (
	"fmt"
	"strings"
) 

func main(){
	// Pendeklarasian Array
	var numbers [4]int
	numbers = [4]int{1, 2, 3, 4}

	var newString = [3]string{"Abdul", "Hanif", "Pratama"}

	fmt.Printf("%#v\n", numbers)
	fmt.Printf("%#v\n", newString)

	fmt.Println("====================================")
	fmt.Println()


	/*
		verb %#v digunakan untuk memformat tipe data agar
		kita dapat melihat panjang arraynya,
	*/


	// Modify Element Through Index
	var fruits = [3]string{"Apel", "Pisang", "Jeruk"}

	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Orange"

	fmt.Printf("%#v\n", fruits)
	fmt.Println("====================================")
	fmt.Println()



	// Inisialisasi Nilai Awal Array tanpa jumlah elemen
	var cars = [...]string{"Ferrari", "Lamborghini"}

	fmt.Println("Jumlah array \t:", len(cars))
	fmt.Println("Isi array \t:", fruits)

	fmt.Println("====================================")
	fmt.Println()


	// loop through elements
	var colors = [3]string{"blue", "green", "yellow"}


	// Cara pertama
	for i, v := range colors {
		fmt.Printf("Index: %d, Value: %s\n", i, v)
	}

	fmt.Println(strings.Repeat("#", 25))

	// Cara kedua
	for i := 0; i < len(colors); i++{
		fmt.Printf("Index: %d, Value: %s\n", i, colors[i])
	}

	fmt.Println("====================================")
	fmt.Println()



	// Multidimensional Array
	balances := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	for _, arr := range balances {
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}

	fmt.Println("====================================")
	fmt.Println()

}