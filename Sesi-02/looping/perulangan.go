package main

import "fmt"

func main() {
	// Perulangan menggunakan keyword for (first way of looping)
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}
	fmt.Println("=====================")
	fmt.Println()

	// Penggunaan Keyword for Dengan Argumen Hanya Kondisi (second way of looping)
	var a = 0

	for a < 5 {
		fmt.Println("Angka", a)
		a++
	}

	fmt.Println("=====================")
	fmt.Println()

	// Penggunaan keyword for tanpa argumen (third way of looping)
	var b = 0
	for {
		fmt.Println("Angka", b)

		b++
		if b == 5 {
			break
		}
	}
	fmt.Println("=====================")
	fmt.Println()

	// Loopings (break and continue keywords)
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}
		if i > 8 {
			break
		}
		fmt.Println("Angka", i)
	}
	fmt.Println("=====================")
	fmt.Println()


	// Nested Looping
	for x := 0; x < 5; x++{
		for y := x; y < 5; y++ {
			fmt.Print(y ," ")
		}

		fmt.Println()
	}

	fmt.Println("=====================")
	fmt.Println()


	// Looping labels
	outerLoop:
	for m := 0; m < 3; m++{
		fmt.Println("Perulangan ke -", m + 1)
		for s := 0; s < 3; s++{
			if m == 2 {
				break outerLoop
			}
			fmt.Print(s ," ")
		}
	}
	fmt.Println()
}
