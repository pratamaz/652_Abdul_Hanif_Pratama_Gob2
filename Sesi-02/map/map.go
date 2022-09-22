package main

import (
	"fmt"
	"strings"
)

func main() {
	var person map[string]string // Deklarasi
	
	// Cara 2	
	person = map[string]string{
		"name"   : "Abdul",
		"age "   : "23",
		"Addres" : "Jalan Sudirman",
	}

	// Cara 2
	// person["name"] = "Abdul"
	// person["age"] = "23"
	// person["addres"] = "Jalan Sudirman"

	// fmt.Println("Nama  :", person["name"])
	// fmt.Println("Age   :", person["age"])
	// fmt.Println("Addres:", person["addres"])

	// Atau bisa dengan looping
	for key, value := range person{
		fmt.Println(key, ":", value)
	}

	fmt.Printf(strings.Repeat("=", 30))
	fmt.Printf(strings.Repeat("\n", 2))


	// Menghapus Value pada Map
	delete(person, "age")

	fmt.Println("After deleting:", person)

	fmt.Printf(strings.Repeat("=", 30))
	fmt.Printf(strings.Repeat("\n", 2))


	// Mendeteksi Value pada Map
	value, exist := person["Addres"]
	
	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value doesn't exist")
	}

	value1, exist1 := person["age"]
	
	if exist1 {
		fmt.Println(value1)
	} else {
		fmt.Println("Value doesn't exist")
	}

	fmt.Printf(strings.Repeat("=", 30))
	fmt.Printf(strings.Repeat("\n", 2))


	// Menggabungkan Slice dengan Map
	var people = []map[string]string{
		{"nama" : "Abdul", "umur": "22"},
		{"nama" : "Hanif", "umur": "21"},
		{"nama" : "Pratama", "umur": "23"},
	}

	for i, person := range people {
		fmt.Printf("Index: %d, nama: %s, age: %s\n", i, person["nama"], person["umur"])
	}

	fmt.Printf(strings.Repeat("=", 30))
	fmt.Printf(strings.Repeat("\n", 2))
}