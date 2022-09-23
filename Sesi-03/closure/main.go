package main

import (
	"fmt"
	"strings"
)

func main() {
	// Deklarasi Closure - 1
	fmt.Println("====================================")
	fmt.Println("         Closure Function 1         ")
	fmt.Println("====================================")

	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	var numbers = []int{2, 3, 4, 3, 4, 2, 3}
	var min, max = getMinMax(numbers)
	fmt.Printf("Data : %v\nmin : %v\nmax : %v\n", numbers, min, max)

	fmt.Println()


	// Deklarasi Closure - 2
	fmt.Println("====================================")
	fmt.Println("         Closure Function 2         ")
	fmt.Println("====================================")

	var evenNumbers = func (numbers1 ...int) []int {
		var result []int

		for _, v := range numbers1 {
			if v % 2 == 0 {
				result = append(result, v)
			}
		}

		return result
	}
	var numbers1 = []int{4, 93, 77, 10, 52, 22, 34}

	fmt.Println(evenNumbers(numbers1...))
	
	fmt.Println()

	
	// Closure (IIFE)
	fmt.Println("====================================")
	fmt.Println("           Closure IIFE             ")
	fmt.Println("====================================")

	var isPalindrome = func (str string) bool  {
		var temp string = ""
		
		for i := len(str) - 1; i >= 0; i-- {
			temp += string(byte(str[i]))
		}
		
		return temp == str
	}("katak")

	fmt.Println(isPalindrome)

	fmt.Println()


	// Closure sebagai Nilai Kembalian
	fmt.Println("====================================")
	fmt.Println("     Closure as a Return Value      ")
	fmt.Println("====================================")

	var studentLists = []string{"Abdul", "Hanif","Pratama", "Patrick", "Krab"}

	var find = findStudent(studentLists)

	fmt.Println(find("Hanif"))

	fmt.Println()


	fmt.Println("====================================")
	fmt.Println("          Closure Callback          ")
	fmt.Println("====================================")

	var numbers2 = []int{2, 5, 8, 10, 3, 99, 23}

	var find1 = findOddNumbers(numbers2, func(number int)bool {
		return number % 2 != 0
	})

	fmt.Println("Total odd numbers:", find1)

	fmt.Println()
}

func findStudent(students []string) func(string) string {
	
	return func(s string) string {
		var student string
		var position int

		for i, v := range students {
			if strings.ToLower(v) == strings.ToLower(s) {
				student = v
				position = i
				break
			}
		}

		if student == "" {
			return fmt.Sprintf("%s does'nt exist!!")
		}
		return fmt.Sprintf("We found %s at position %d", s, position + 1)
	}
}


func findOddNumbers(numbers []int, callback func(int) bool) int {
	var totalOddNumbers int

	for _, v := range numbers {
		if callback(v) {
			totalOddNumbers++
		}
	}
	return totalOddNumbers
}