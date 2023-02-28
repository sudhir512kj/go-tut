package main

import "fmt"

// Author structure
type author struct {
	name      string
	branch    string
	particles int
	salary    int
}

type data int

// Defining a method with
// non-struct type receiver
func (d1 data) multiply(d2 data) data {
	return d1 * d2
}

// Method with a receiver
// of author type
func (a author) show1() {
	a.branch = "Rating"
	fmt.Println("Author's Name: ", a.name)
	fmt.Println("Branch Name: ", a.branch)
	fmt.Println("Published articles: ", a.particles)
	fmt.Println("Salary: ", a.salary)
}

// Method with a receiver of author type
func (a *author) show2(abranch string) {
	(*a).branch = abranch
}

// Creating structures
type student struct {
	name   string
	branch string
}

type teacher struct {
	language string
	marks    int
}

// Same name methods, but with
// different type of receivers
func (s student) show() {

	fmt.Println("Name of the Student:", s.name)
	fmt.Println("Branch: ", s.branch)
}

func (t teacher) show() {

	fmt.Println("Language:", t.language)
	fmt.Println("Student Marks: ", t.marks)
}

// Main function
func main() {

	// Initializing the values
	// of the author structure
	res := author{
		name:      "Sona",
		branch:    "CSE",
		particles: 203,
		salary:    34000,
	}

	// Calling the method
	res.show1()
	fmt.Println(res.branch)

	value1 := data(23)
	value2 := data(20)
	res2 := value1.multiply(value2)
	fmt.Println("Final result: ", res2)

	res3 := author{
		name:   "Sona",
		branch: "CSE",
	}

	fmt.Println("Author's name: ", res3.name)
	fmt.Println("Branch Name(Before): ", res3.branch)

	// Creating a pointer
	p := &res3

	// Calling the show method
	p.show2("ECE")
	fmt.Println("Author's name: ", res3.name)
	fmt.Println("Branch Name(After): ", res3.branch)

	// Initializing values
	// of the structures
	val1 := student{"Rohit", "EEE"}

	val2 := teacher{"Java", 50}

	// Calling the methods
	val1.show()
	val2.show()
}
