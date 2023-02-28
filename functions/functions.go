package main

import (
	"fmt"
	"strings"
)

// variadic functions
func joinstr(elements ...string) string {
	return strings.Join(elements, "-")
}

// passing anonymous func as an argument
func GFG(i func(p, q string) string) {
	fmt.Println(i("Geeks", "for"))
}

// return anonymous func
func GFG2() func(i, j string) string {
	myf := func(i, j string) string {
		return i + j + "geeksforgeeks"
	}
	return myf
}

func modify(Z int) {
	Z = 70
}

func modifyPointer(Z *int) {
	*Z = 70
}

func swap(x, y int) int {
	var temp int
	temp = x
	x = y
	y = temp

	return temp
}

func swapPointer(x, y *int) int {

	// taking a temporary variable
	var tmp int

	tmp = *x
	*x = *y
	*y = tmp

	return tmp
}

// return multiple values
func myfunc(p, q int) (int, int, int) {
	return p - q, p * q, p + q
}

func myfunc2(p, q int) (rectange int, square int) {
	rectange = p * q
	square = p * q
	return
}

func main() {

	// anonymous function
	func() {
		fmt.Println("Welcome to Geeksforgeeks")
	}()

	// anonymous function with a name
	anoFunc := func() {
		fmt.Println("Welcome to Geeksforgeeks with a name")
	}
	anoFunc()

	// anonymous func with arguments
	func(ele string) {
		fmt.Println(ele)
	}("Geeksforgeeks")

	anoFunc2 := func(p, q string) string {
		return p + q + "Geeks"
	}
	// passing anonymous func as an argument
	GFG(anoFunc2)

	value := GFG2()
	fmt.Println(value("Welcome", "to"))

	fmt.Println(joinstr())
	fmt.Println(joinstr("GEEK", "GFG"))
	fmt.Println(joinstr("Geeks", "for", "Geeks"))
	fmt.Println(joinstr("G", "E", "E", "k", "S"))
	elements := []string{"geeks", "FOR", "geeks"}
	fmt.Println(joinstr(elements...))

	var Z int = 10
	fmt.Printf("Before function call, value of Z is : %d", Z)

	modify(Z)
	fmt.Printf("\nAfter function call, value of Z is : %d", Z)

	modifyPointer(&Z)
	fmt.Printf("\nAfter function call, value of Z is : %d", Z)

	var f int = 700
	var s int = 900

	fmt.Printf("Values Before Function Call\n")
	fmt.Printf("f = %d and s = %d\n", f, s)

	// call by values
	swap(f, s)

	fmt.Printf("\nValues After Function Call\n")
	fmt.Printf("f = %d and s = %d", f, s)

	// call by reference
	swapPointer(&f, &s)

	fmt.Printf("\nValues After Function Call\n")
	fmt.Printf("f = %d and s = %d", f, s)

	// The return values are assigned into
	// three different variables
	var myvar1, myvar2, myvar3 = myfunc(4, 2)

	// Display the values
	fmt.Printf("Result is: %d", myvar1)
	fmt.Printf("\nResult is: %d", myvar2)
	fmt.Printf("\nResult is: %d", myvar3)

	// The return values are assigned into
	// two different variables
	var area1, area2 = myfunc2(2, 4)

	// Display the values
	fmt.Printf("Area of the rectangle is: %d", area1)
	fmt.Printf("\nArea of the square is: %d", area2)
}
