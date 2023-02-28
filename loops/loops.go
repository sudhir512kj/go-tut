package main

import "fmt"

func main() {
	for i := 0; i < 4; i++ {
		fmt.Println("GeeksforGeeks")
	}

	// Infinite loop
	// for {
	// 	fmt.Printf("GeeksforGeeks\n")
	//   }

	i := 0
	for i < 3 {
		i += 2
	}
	fmt.Println(i)

	// Here rvariable is a array
	rvariable := []string{"GFG", "Geeks", "GeeksforGeeks"}

	// i and j stores the value of rvariable
	// i store index number of individual string and
	// j store individual string of the given array
	for i, j := range rvariable {
		fmt.Println(i, j)
	}

	// String as a range in the for loop
	for i, j := range "XabCd" {
		fmt.Printf("The index number of %U is %d\n", j, i)
	}

	// using maps
	mmap := map[int]string{
		22: "Geeks",
		33: "GFG",
		44: "GeeksforGeeks",
	}
	for key, value := range mmap {
		fmt.Println(key, value)
	}

	// using channel
	chnl := make(chan int)
	go func() {
		chnl <- 100
		chnl <- 1000
		chnl <- 10000
		chnl <- 100000
		close(chnl)
	}()
	for i := range chnl {
		fmt.Println(i)
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
		// For loop breaks when the value of i = 3
		if i == 3 {
			break
		}
	}

	var x int = 0

	// for loop work as a while loop
Lable1:
	for x < 8 {
		if x == 5 {

			// using goto statement
			x = x + 1
			goto Lable1
		}
		fmt.Printf("value is: %d\n", x)
		x++
	}

	var y int = 0

	// for loop work as a while loop
	for y < 8 {
		if y == 5 {

			// skip two iterations
			y = y + 2
			continue
		}
		fmt.Printf("value is: %d\n", y)
		y++
	}

	switch day := 4; day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid")
	}

	var value string = "five"

	// Switch statement without default statement
	// Multiple values in case statement
	switch value {
	case "one":
		fmt.Println("C#")
	case "two", "three":
		fmt.Println("Go")
	case "four", "five", "six":
		fmt.Println("Java")
	}

	var value2 interface{}
	switch q := value2.(type) {
	case bool:
		fmt.Println("value is of boolean type")
	case float64:
		fmt.Println("value is of float64 type")
	case int:
		fmt.Println("value is of int type")
	default:
		fmt.Printf("value is of type: %T", q)

	}

	// Creating a channel
	// Here  deadlock arises because
	// no goroutine is writing
	// to this channel so, the select
	// statement has been blocked forever
	// c := make(chan int)
	// select {
	// case <-c:
	// }

	// Creating a channel
	c := make(chan int)
	select {
	case <-c:
	default:
		fmt.Println("!.. Default case..!")
	}
}
