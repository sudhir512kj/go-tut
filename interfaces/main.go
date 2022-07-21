// Golang program illustrates how
// to implement an interface
package main

import "fmt"

// Creating an interface
type tank interface {

	// Methods
	Tarea() float64
	Volume() float64
}

type myvalue struct {
	radius float64
	height float64
}

// Implementing methods of
// the tank interface
func (m myvalue) Tarea() float64 {

	return 2*m.radius*m.height +
		2*3.14*m.radius*m.radius
}

func (m myvalue) Volume() float64 {

	return 3.14 * m.radius * m.radius * m.height
}

func myfun(a interface{}) {

	// Extracting the value of a
	val, ok := a.(string)
	fmt.Println("Value: ", val)

	// invalid case
	value, ok := a.(float64)
	fmt.Println(value, ok)
}

// Interface 1
type AuthorDetails interface {
	details()
}

// Interface 2
type AuthorArticles interface {
	articles()
}

// Structure
type author struct {
	a_name    string
	branch    string
	college   string
	year      int
	salary    int
	particles int
	tarticles int
}

// Implementing method
// of the interface 1
func (a author) details() {

	fmt.Printf("Author Name: %s", a.a_name)
	fmt.Printf("\nBranch: %s and passing year: %d", a.branch, a.year)
	fmt.Printf("\nCollege Name: %s", a.college)
	fmt.Printf("\nSalary: %d", a.salary)
	fmt.Printf("\nPublished articles: %d", a.particles)

}

// Implementing method
// of the interface 2
func (a author) articles() {

	pendingarticles := a.tarticles - a.particles
	fmt.Printf("\nPending articles: %d", pendingarticles)
}

/*
Embedding Interfaces
*/
type FinalDetails interface {
	AuthorArticles
	AuthorDetails
}

/*
Polymorphism
*/

type employee interface {
	develop() int
	name() string
}

// Structure 1
type team1 struct {
	totalapp_1 int
	name_1     string
}

// Methods of employee interface are
// implemented by the team1 structure
func (t1 team1) develop() int {
	return t1.totalapp_1
}

func (t1 team1) name() string {
	return t1.name_1
}

// Structure 2
type team2 struct {
	totalapp_2 int
	name_2     string
}

// Methods of employee interface are
// implemented by the team2 structure
func (t2 team2) develop() int {
	return t2.totalapp_2
}

func (t2 team2) name() string {
	return t2.name_2
}

func finaldevelop(i []employee) {

	totalproject := 0
	for _, ele := range i {

		fmt.Printf("\nProject environment = %s\n ", ele.name())
		fmt.Printf("Total number of project %d\n ", ele.develop())
		totalproject += ele.develop()
	}
	fmt.Printf("\nTotal projects completed by "+
		"the company = %d", totalproject)
}

// Main Method
func main() {

	// Accessing elements of
	// the tank interface
	var t tank
	t = myvalue{10, 14}
	fmt.Println("Area of tank :", t.Tarea())
	fmt.Println("Volume of tank:", t.Volume())

	var val interface {
	} = "GeeksforGeeks"

	myfun(val)

	var a1 interface {
	} = 98.09

	myfun(a1)

	var a2 interface {
	} = "GeeksforGeeks"

	myfun(a2)

	values := author{
		a_name:    "Mickey",
		branch:    "Computer science",
		college:   "XYZ",
		year:      2012,
		salary:    50000,
		particles: 209,
		tarticles: 309,
	}

	// Accessing the method
	// of the interface 1
	var i1 AuthorDetails = values
	i1.details()

	// Accessing the method
	// of the interface 2
	var i2 AuthorArticles = values
	i2.articles()

	// Embedded Interface FinalDetails
	values = author{
		a_name:    "Sudhir Meena",
		branch:    "Computer science",
		college:   "XYZ",
		year:      2012,
		salary:    50000,
		particles: 209,
		tarticles: 309,
	}

	// Accessing the methods of
	// the interface 1 and 2
	// Using FinalDetails interface
	var f FinalDetails = values
	f.details()
	f.articles()

	res1 := team1{totalapp_1: 20,
		name_1: "Android"}

	res2 := team2{totalapp_2: 35,
		name_2: "IOS"}

	final := []employee{res1, res2}
	finaldevelop(final)
}
