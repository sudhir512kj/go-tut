package main

import (
	"fmt"
	"time"
)

func display(str string) {
	for w := 0; w < 6; w++ {
		time.Sleep(1 * time.Second)
		fmt.Println(str)
	}
}

// function 1
func portal1(channel1 chan string) {

	time.Sleep(10 * time.Second)
	channel1 <- "Welcome to channel 1"
}

// function 2
func portal2(channel2 chan string) {

	time.Sleep(9 * time.Second)
	channel2 <- "Welcome to channel 2"
}

// function 1
func portal3(channel1 chan string) {
	for i := 0; i <= 3; i++ {
		channel1 <- "Welcome to channel 1"
	}

}

// function 2
func portal4(channel2 chan string) {
	channel2 <- "Welcome to channel 2"
}

// For goroutine 1
func Aname() {

	arr1 := [4]string{"Rohit", "Suman", "Aman", "Ria"}

	for t1 := 0; t1 <= 3; t1++ {

		time.Sleep(150 * time.Millisecond)
		fmt.Printf("%s\n", arr1[t1])
	}
}

// For goroutine 2
func Aid() {

	arr2 := [4]int{300, 301, 302, 303}

	for t2 := 0; t2 <= 3; t2++ {

		time.Sleep(500 * time.Millisecond)
		fmt.Printf("%d\n", arr2[t2])
	}
}

func main() {

	// Calling Goroutine
	go display("Welcome")

	// Calling normal function
	display("GeeksforGeeks")

	fmt.Println("Welcome!! to Main function")

	// Creating Anonymous Goroutine
	go func() {

		fmt.Println("Welcome!! to GeeksforGeeks")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("GoodBye!! to Main function")

	// Creating channels
	R1 := make(chan string)
	R2 := make(chan string)

	// calling function 1 and
	// function 2 in goroutine
	go portal1(R1)
	go portal2(R2)

	select {
	// case 1 for portal 1
	case op1 := <-R1:
		fmt.Println(op1)

	// case 2 for portal 2
	case op2 := <-R2:
		fmt.Println(op2)
	}

	// Creating channels
	R3 := make(chan string)
	R4 := make(chan string)

	// calling function 1 and
	// function 2 in goroutine
	go portal1(R3)
	go portal2(R4)

	// the choice of selection
	// of case is random
	select {
	case op1 := <-R3:
		fmt.Println(op1)
	case op2 := <-R4:
		fmt.Println(op2)
	}

	fmt.Println("!...Main Go-routine Start...!")

	// calling Goroutine 1
	go Aname()

	// calling Goroutine 2
	go Aid()

	time.Sleep(3500 * time.Millisecond)
	fmt.Println("\n!...Main Go-routine End...!")
}
