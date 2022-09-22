package main

import (
	"fmt"
	"strings"
)

func main() {
	// Deklarasi variabel dengan tipe data uint8
	var a uint8 = 10
	var b byte // byte adalah alias dari tipe data uint8

	b = a // tidak error

	fmt.Println(b == a)

	fmt.Printf(strings.Repeat("=", 30))
	fmt.Printf(strings.Repeat("\n", 2))

	// Mendeklarasikan tipe data alias bernama second
	type second = uint

	var hour second = 3600
	fmt.Printf("Hour type: %T\n", hour)

	fmt.Printf(strings.Repeat("=", 30))
	fmt.Printf(strings.Repeat("\n", 2))
}
