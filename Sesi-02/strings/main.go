package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
) 

func main() {
	// Looping Over String (byte-by-byte)
	city := "Jakarta"

	for i := 0; i < len(city); i++ {
		fmt.Printf("Index: %d, byte: %d\n", i, city[i])
	}

	fmt.Printf(strings.Repeat("=", 30))
	fmt.Printf(strings.Repeat("\n", 2))


	// Looping Over String (Run-by-run)
	str1 := "Makan"
	str2 := "mânca"

	// Mencari Jumlah Byte
	fmt.Printf("Str1 byte length => %d\n", len(str1))
	fmt.Printf("Str2 byte length => %d\n", len(str2))

	// Mencari Jumlah Character
	fmt.Printf("Str1 character length => %d\n", utf8.RuneCountInString(str1))
	fmt.Printf("Str2 character length => %d\n", utf8.RuneCountInString(str2))
	
	fmt.Printf(strings.Repeat("=", 30))
	fmt.Printf(strings.Repeat("\n", 2))
}