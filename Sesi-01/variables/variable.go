package main

import "fmt"

func main() {
	// Belajar Membuat Variabel
	fmt.Println("===================================")
	fmt.Println("        Belajar Variables          ")
	fmt.Println("===================================")

	// Variabel Dengan Tipe Data
	var namaLengkap string = "Abdul"
	var umur int = 21

	fmt.Println("     Variabel Dengan Tipe Data     ")
	fmt.Printf("\n")

	fmt.Println("Ini adalah namanya ==>", namaLengkap)
	fmt.Println("Ini adalah umurnya ==>", umur)
	fmt.Println("===================================")
	fmt.Printf("\n")

	// Variabel Tanpa Tipe Data
	var namaAnda = "Abdul"
	var umurAnda = 21

	fmt.Println("     Variabel Tanpa Tipe Data     ")
	fmt.Printf("\n")

	fmt.Println("Ini adalah namanya ==>", namaAnda)
	fmt.Println("Ini adalah umurnya ==>", umurAnda)
	fmt.Printf("%T %T\n", namaAnda, umurAnda)
	fmt.Println("===================================")
	fmt.Printf("\n")

	// Variabel Tanpa Tipe Data (Short Declaration)
	namaAku := "Hanif"
	umurAku := 21
	
	fmt.Println("     Variabel Short Declaration    ")
	fmt.Printf("\n")

	fmt.Printf("%T, %T\n", namaAku, umurAku)
	fmt.Println("===================================")
	fmt.Printf("\n")

	// Multivariable
	var one, two, three string = "1", "2", "3"
	var angkaOne, angkaTwo, angkaThree int = 1, 2, 3

	fmt.Println("     Multivariable Declaration     ")
	fmt.Printf("\n")

	fmt.Println("Test Data String Multiple :", one, two, three)
	fmt.Println("Test Data Integer Multiple:", angkaOne, angkaTwo, angkaThree)
	fmt.Println("===================================")
	fmt.Printf("\n")

	// Multivariable Short Declaration
	var oneData, twoData, threeData = 1, true, "Abdul Hanif"
	dataOne, dataTwo, dataThree := true, 2.2, "Pratama"

	fmt.Println("   Multivariable Short Declaration  ")
	fmt.Printf("\n")

	fmt.Println("Test short multiple", oneData, twoData, threeData)
	fmt.Println("Test short multiple", dataOne, dataTwo, dataThree)
	fmt.Println("===================================")
	fmt.Printf("\n")

	// Underscore Variable
	fmt.Println("         Underscore Variabel       ")
	fmt.Printf("\n")

	var testVariabel string
	var dataNama, dataUmur, dataAlamat = "Abdul", 21, "Bekasi"
	_, _, _, _ = testVariabel, dataNama, dataUmur, dataAlamat

	fmt.Printf("Test Underscore Variabel > %T\n", testVariabel)
	fmt.Printf("Halo, nama saya %s, umur saya %d, saya beralamat di %s\n", 
				dataNama, dataUmur, dataAlamat)
	fmt.Println("===================================")
	fmt.Printf("\n")

}