package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// Membuat goroutine
	// go goroutine()

	fmt.Println("Main Execution Started")

	go firstProcess(8)
	secondProcess(8)

	fmt.Println("No. of Goroutines:", runtime.NumGoroutine())

	time.Sleep(time.Second * 2)

	fmt.Println("Main Execution Ended")
}

// func goroutine() {
// 	fmt.Println("Hello")
// }

// Asycronous Process
func firstProcess(index int) {
	fmt.Println("First Process Func Started")
	for i := 1; i <= index; i++ {
		fmt.Println("i =", i)
	}
	fmt.Println("First Process Func Ended")
}

func secondProcess(index int) {
	fmt.Println("Second Process Func Started")
	for j := 1; j <= index; j++ {
		fmt.Println("j =", j)
	}
	fmt.Println("Second Process Func Ended")
}