package main

import (
	"fmt"
)

func main(){
	fmt.Println("===================================")
	fmt.Println("        Belajar Hello World        ")
	fmt.Println("===================================")
	// Belajar Hello World
	fmt.Println("Hello World!")

	// Mencetak dengan fmt.Printf
	fmt.Printf("Abdul", "Belajar", "Koding")

	// mencetak dengan fmt.Println
	fmt.Println("Hello")
	fmt.Println("Hello", "Abdul")
	fmt.Println("Abdul", "Belajar", "Koding")

	// Mencetak dengan fmt.Print
	fmt.Print("Abdul", "Belajar", "Koding")
	fmt.Print("Hello Abdul\n")
	fmt.Print("Abdul", "Hanif\n")
	fmt.Print("Abdul", " ", "Hanif\n")
	fmt.Println("===================================")
	fmt.Printf("\n\n")
		
	// Komentar Pada Golang
	fmt.Println("===================================")
	fmt.Println("        Belajar Komentar           ")
	fmt.Println("===================================")
	
	fmt.Println(`"/**/", Komentar ini digunakan
				untuk mengomentari lebih dari 1
				baris	
	`)
	/*
		komentar Untuk Lebih dari
		Satu baris
	*/
	fmt.Println(`"//", Komentar ini digunakan
				untuk mengomentari hanya 1 baris	
	`)

	// fmt.Println("Baris ini tidak akan dieksekusi")


	fmt.Println("===================================")
	fmt.Printf("\n\n")
}