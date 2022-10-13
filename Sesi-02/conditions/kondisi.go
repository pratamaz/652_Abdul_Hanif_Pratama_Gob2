package main

import "fmt"

func main(){
	// Belajar Conditions (if, else if, dan else)
	fmt.Println("===================================")
	fmt.Println("       Belajar Conditions 1        ")
	fmt.Println()

	var point = 1
	
	if point == 10 {
		fmt.Println("Lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("Lulus")
	} else if point == 4 {
		fmt.Println("Hampir Lulus")
	} else {
		fmt.Printf("Tidak Lulus. Nilai Anda %d\n", point)
	}
	fmt.Println()

	
	// Belajar Conditions (Temporary Variable)
	fmt.Println("===================================")
	fmt.Println("       Belajar Conditions 2        ")
	fmt.Println()

	var currentYear = 2010

	if age := currentYear - 1998; age < 17 {
		fmt.Printf("Kamu belum boleh membuat SIM ==> %d", age)
	}else {
		fmt.Println("Kamu sudah boleh membuat SIM")
	}

	fmt.Println()

	/*
		Pada kodingan di atas terdapat variabel currentYear
		yang bernilai 2022 (tahun sekarang). 
		Lalu kita buat kondisi dengan membuat variabel baru 
		yaitu variabel age yang dimana nilainya hasil dari
		kalkulasi nilai variabel currentYear dikurangi dengan
		nilai 1998, dan kondisi kedua yaitu jika nilai dari var
		age kurang dari 17 maka belum boleh membuat SIM.
	*/


	// Belajar Switch Case
	fmt.Println("===================================")
	fmt.Println("          Belajar Switch           ")
	fmt.Println()
	
	var score = 8

	switch score{
	case 8:
		fmt.Println("Perfect")
	case 7, 6, 5:
		fmt.Println("Awesome")
	default :
		fmt.Println("Not Bad")		
	}

	fmt.Println("===================================")
	fmt.Println()


	// Switch with relational operators
	var score1 = 6 


	fmt.Println("  Switch with Relational Operators  ")
	fmt.Printf("\n")

	switch {
	case score1 == 8:
		fmt.Println("Perfect")
	case (score1 < 8) && (score1 > 3):
		fmt.Println("Not Bad")
	default :
		{
			fmt.Println("Study Harder")
			fmt.Println("You need to learn more")
		}		
	}

	fmt.Println("===================================")
	fmt.Println()


	// Switch fallthrough keyword
	var score2 = 6 


	fmt.Println(" Switch with (Fallthrough Keyword) ")
	fmt.Printf("\n")

	switch {
	case score2 == 8:
		fmt.Println("Perfect")
	case (score2 < 8) && (score2 > 3):
		fmt.Println("Not Bad")
		fallthrough
	case score2 < 5 :
		fmt.Println("It's ok, but please study harder")	
	default :
		{
			fmt.Println("Study Harder")
			fmt.Println("You need to learn more")
		}		
	}

	fmt.Println("===================================")
	fmt.Println()

	/*
		keyword fallthrough digunakan untuk melanjutkan
		pengecekan kepada case selanjutnya walaupun case
		telah terpenuhi kondisinya. 
	*/


	// Nested Conditions
	var score3 = 10
	
	fmt.Println("         Nested Conditions        ")
	fmt.Printf("\n")	


	if score3 > 7 {
		switch score3 {
		case 10:
			fmt.Println("Perfect")
		default:
			fmt.Println("Nice")	
		}
	}else {
		if score3 == 5 {
			fmt.Println("Not Bad")
		} else if score3 == 3 {
			fmt.Println("Keep trying")
		} else{
			fmt.Println("You can do it!")
			if score3 == 0 {
				fmt.Println("Try harder")
			}
		}
	}

	fmt.Println("===================================")
	fmt.Println()

}