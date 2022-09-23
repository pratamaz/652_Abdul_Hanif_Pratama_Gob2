package main

import (
	"fmt"
	"math"
	"strings"
)

// Function Main
func main() {
	fmt.Println("====================================")
	fmt.Println("             Function               ")
	fmt.Println("====================================")

	
	greet("Abdul", 22)

	
	fmt.Println("====================================")
	fmt.Println()

	fmt.Println("====================================")
	fmt.Println("  Function with Multiple Parameter  ")
	fmt.Println("====================================")

	
	greetDetail("Hanif", "Bekasi")


	fmt.Println("====================================")
	fmt.Println()

	fmt.Println("====================================")
	fmt.Println("           Return Function          ")
	fmt.Println("====================================")

	var names = []string{"Abdul", "Hanif"}
	var printMsg = greetReturn("Heii", names)

	fmt.Println(printMsg)

	fmt.Println("====================================")
	fmt.Println()

	fmt.Println("====================================")
	fmt.Println("Return Function with Multiple Values")
	fmt.Println("====================================")

	var diameter float64 = 15
	var area, circumference = calculate(diameter)
	fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
	fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)

	fmt.Println("====================================")
	fmt.Println()

	fmt.Println("====================================")
	fmt.Println("         Variadic Function 1        ")
	fmt.Println("====================================")

	studentLists := print("Abdul", "Hanif", "Pratama", "Ujang", "Udin")

	for a := range studentLists{
		fmt.Printf("%v\n", studentLists[a])
	}

	fmt.Println("====================================")
	fmt.Println()


	fmt.Println("====================================")
	fmt.Println("         Variadic Function 2        ")
	fmt.Println("====================================")

	numberLists := []int{1, 2, 3, 4, 5, 6, 7, 8}

	result := sum(numberLists...)

	fmt.Println("Result:", result)

	fmt.Println("====================================")
	fmt.Println()


	fmt.Println("====================================")
	fmt.Println("         Variadic Function 3        ")
	fmt.Println("====================================")

	profile("Abdul", "Sate", "Bakso", "Mie Ayam", "Nasi Goreng")

	fmt.Println("====================================")
	fmt.Println()
}

// Membuat greet
func greet(name string, age int8) {
	fmt.Printf("Hello there! My name is %s and I'm %d years old\n", name, age)
}

// Membuat function greetDetail
func greetDetail(nama, alamat string) {
	fmt.Println("Hello there! My name is", nama)
	fmt.Println("I live in", alamat)
}

// membuat function return
func greetReturn(msg string, names []string) string {
	var joinStr = strings.Join(names, " ")

	var result string = fmt.Sprintf("%s %s", msg, joinStr)

	return result

	/*
		Function Sprintf memiliki fungsi yang sama seperti
		function Printf,namun bedanya adalah function Sprintf
		akan me-return sebuah nilaisedang function Printf tidak.
	*/
}

// Returning Multiple Values
func calculate(d float64) (float64, float64) {
	// menghitung luas
	var area float64 = math.Pi * math.Pow(d/2, 2)

	// menghitung keliling
	var circumference = math.Pi * d
	
	// kembalikan 2 nilai
	return area, circumference
}

// Predifined Return Value
// func calculate(d float64) (area float64, circumference float64) {
// 	// hitung luas
// 	area = math.Pi * math.Pow(d / 2, 2)
// 	// hitung keliling
// 	circumference = math.Pi * d

// 	// kembalikan 2 nilai
// 	return
// }

func print(names1 ...string) []map[string]string {
	var result []map[string]string

	for i, v := range names1 {
		key := fmt.Sprintf("Student %d", i+1)
		temp := map[string]string{
			key: v,
		}
		result = append(result, temp)
	}

	return result
}

// Variadic Function 2
func sum(numbers ...int)int {
	total := 0

	for _, v := range numbers {
		total += v
	}
	return total
}

func profile(namaSaya string, favFoods ...string) {
	mergeFavFoods := strings.Join(favFoods, ",")

	fmt.Println("Hello there!! I'm", namaSaya)
	fmt.Println("I really love to eat", mergeFavFoods)
}