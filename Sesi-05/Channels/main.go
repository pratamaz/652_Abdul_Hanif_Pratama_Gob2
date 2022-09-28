package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// Mendeklarasikan Channels
	c := make(chan string)

	d := make(chan string)

	// Mengirim data ke channel
	// c <- value

	// Menerima data dari channel
	// result := <- c

	// Implementing Channel
	fmt.Println("         Implementing Channel        ")

	go introduce("Abdul", c)
	go introduce("Hanif", c)
	go introduce("Pratama", c)

	msg1 := <-c
	fmt.Println(msg1)
	msg2 := <-c
	fmt.Println(msg2)
	msg3 := <-c
	fmt.Println(msg3)

	close(c)
	/*
		function close digunakan untuk mengindikasikan bahwa
		sebuah channel sudah tidak digunakan untuk berkomunikasi lagi.
	*/

	fmt.Println(strings.Repeat("=", 45))
	fmt.Println()



	// Channel with Anonymous function
	fmt.Println("   Channel with Anonymous Function   ")
	fmt.Println()

	students := []string{"Abdul", "Hanif", "Pratama"}

	for _, v := range students {
		go func(student string) {
			fmt.Println("Student", student)
			result1 := fmt.Sprintf("Hai, my name is %s", student)
			d <- result1
		}(v)
	}
	for i := 1; i <= 3; i++ {
		print(d)
	}
	close(d)

	fmt.Println(strings.Repeat("=", 45))
	fmt.Println("")



	// Channel Directions
	fmt.Println("         Channel Directions       ")
	fmt.Println()

	e := make(chan string)

	people := []string{"Udin", "Ujang", "Mamad"}

	for _, v := range people {
		go introduce1(v, e)
	}

	for i := 1; i <= 3; i++ {
		print2(e)
	}

	fmt.Println(strings.Repeat("=", 45))
	fmt.Println("")



	// Unbuffered Channel vs Buffered Channel
	fmt.Println("        Unbuffered Channel        ")
	fmt.Println()

	c1 := make(chan int)

	go func(f chan int) {
		fmt.Println("Func Goroutine Starts Sending Data Into the Channel")

		f <- 10

		fmt.Println("Func Goroutine Afte Sending Data Into the Channel")
	}(c1)

	fmt.Println("Main Goroutine Sleeps for 2 Seconds")
	time.Sleep(time.Second * 2)

	fmt.Println("Main Goroutine Starts Receiving Data")

	g := <-c1

	fmt.Println("Main Goroutine Received Data:", g)

	close(c1)
	time.Sleep(time.Second)

	fmt.Println(strings.Repeat("=", 45))
	fmt.Println("")

	
	// Buffered Channel
	fmt.Println("         Buffered Channel        ")
	fmt.Println()

	c2 := make(chan int, 3)

	go func(h chan int) {
		for i := 1; i <= 5; i++ {
			fmt.Printf("Func Goroutine %d Starts Sending Data into the Channels\n", i)

			h <- i

			fmt.Printf("Func Goroutine %d After Sending Data into the Channels\n", i)
		}
		close(h)

		}(c2)
		
		fmt.Println("Main Goroutine sleeps 2 seconds")
		time.Sleep(time.Second * 2)
		
		for x := range c2 {
			fmt.Println("Main Goroutine Received Value From Channel:", x)
		}
		
		fmt.Println(strings.Repeat("=", 45))
		fmt.Println("")
		
		
		
	// Channel Select
	fmt.Println("          Channel Select       ")
	fmt.Println()

	d1 := make(chan string)
	d2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)

		d1 <- "Hello!"
	}()

	go func() {
		time.Sleep(1 * time.Second)

		d2 <- "Mantap!"
	}()

	for i := 1; i <= 2; i++ {
		select {
		case pesan1 := <-d1:
			fmt.Println("Received", pesan1)
		case pesan2 := <-d2:
			fmt.Println("Received", pesan2)
		}
	}

	fmt.Println(strings.Repeat("=", 45))
	fmt.Println("")
}



func introduce(student string, c chan string) {
	fmt.Println()

	result := fmt.Sprintf("Hai, my name is %s", student)

	c <- result
}


func print(d chan string) {
	fmt.Println(<-d)
}


func print2(e <-chan string) {
	fmt.Println(<-e)
}


func introduce1(person string, e chan<- string) {
	result3 := fmt.Sprintf("Hai, my name is %s", person)
	e <- result3
}