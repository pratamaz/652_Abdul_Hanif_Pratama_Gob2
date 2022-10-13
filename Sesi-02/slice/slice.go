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



	// Fungsi Copy()
	var fruits1 = []string{"Apple", "Banana", "Mango"}
	var fruits2 = []string{"Durian", "Pineapple", "Starfruit"}

	nn := copy(fruits1, fruits2)

	fmt.Println("Fruits1 =>", fruits1)
	fmt.Println("Fruits2 =>", fruits2)
	fmt.Println("Copied Elements =>", nn)

	fmt.Println("====================================")
	fmt.Println()


	// Slicing Pada Slice
	var buah1 = []string{"Apel", "Pisang", "Durian", "Nanas"}
	
	var buah2 = buah1[1:4]
	fmt.Printf("%#v\n", buah2)

	var buah3 = buah1[0:]
	fmt.Printf("%#v\n", buah3)

	var buah4 = buah1[:3]
	fmt.Printf("%#v\n", buah4)

	var buah5 = buah1[:]
	fmt.Printf("%#v\n", buah5)

	fmt.Println("====================================")
	fmt.Println()


	// Kombinasi slicing dan append
	fruits1 = append(fruits1[1:3], "Avocado")

	fmt.Printf("%#v\n", fruits1)

	fmt.Println("====================================")
	fmt.Println()


	// Backing Array
	var fruits3 = fruits1[2:4]

	fruits3[0] = "Melon"

	fmt.Println("Fruits1 => ", fruits1)
	fmt.Println("Fruits3 => ", fruits3)

	fmt.Println("====================================")
	fmt.Println()


	// Membuat Backing Array
	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Nissan"

	fmt.Println("Cars:", cars)
	fmt.Println("newCars:", newCars)

	fmt.Println("====================================")
	fmt.Println()
}