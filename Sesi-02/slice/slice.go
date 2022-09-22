package main

import "fmt"

func main() {
	// Inisialisasi Slice
	var colors = []string{"Red", "Yellow", "Blue"}
	var newColor = colors[0:2]
	fmt.Println(newColor)

	fmt.Println("====================================")
	fmt.Println()

	// Membuat slice dengan fungsi make
	var fruits = make([]string, 3)

	_ = fruits
	fmt.Printf("%#v\n", fruits)

	fmt.Println("====================================")
	fmt.Println()


	// Menambahkan element pada Slice
	// A. Dengan cara mengakses indexnya
	fruits[0] = "apple"
	fruits[1] = "orange"
	fruits[2] = "mango"

	fmt.Printf("%#v\n", fruits)
	
	fmt.Println("====================================")
	fmt.Println()


	// B. Dengan menggunakan fungsi Append()
	fruits = append(fruits, "banana", "jackfruit")

	fmt.Printf("%#v\n", fruits)

	fmt.Println("====================================")
	fmt.Println()

}