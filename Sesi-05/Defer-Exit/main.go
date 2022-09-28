package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	// Defer #1
	fmt.Println("       Defer #1       ")
	fmt.Println()

	defer fmt.Println("Defer function starts to execute")
	fmt.Println("Hai Everyone")
	fmt.Println("Welcome back to Go learning center")

	fmt.Println(strings.Repeat("=", 35))

	fmt.Println("       Defer #2       ")
	fmt.Println()

	callDeferFunc()
	fmt.Println("Hai Everyone!")

	fmt.Println(strings.Repeat("=", 35))
	fmt.Println()


	fmt.Println("       Exit       ")
	fmt.Println()

	defer fmt.Println("Invoke with Defer")
	fmt.Println("Before Exiting")
	
	fmt.Println(strings.Repeat("=", 35))
	fmt.Println()

	os.Exit(1)
}

func callDeferFunc() {
	defer deferFunc()
}

func deferFunc() {
	fmt.Println("Defer Func Starts to Execute")
}