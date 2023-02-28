package main

import (
	"fmt"
	"reflect"
)

type Address struct {
	Name    string
	City    string
	Pincode int
}

type Car struct {
	Name, Model, Color string
	WeightInKg         float64
}

type Employee struct {
	firstName, lastName string
	age, salary         int
}

// type Author struct {
// 	name   string
// 	branch string
// 	year   int
// }

// type HR struct {
// 	details Author
// }

// Creating structure
type Student struct {
	name   string
	branch string
	year   int
}

// Creating nested structure
type Teacher struct {
	name    string
	subject string
	exp     int
	details Student
}

// Structure
type details struct {

	// Fields of the
	// details structure
	name    string
	age     int
	gender  string
	psalary int
}

// With promoted fields
// Nested structure
type student struct {
	branch string
	year   int
	details
}

// Method 1
func (e employee2) promotmethod(tarticle int) int {
	return e.particle * tarticle
}

// Nested structure
type employee2 struct {
	post     string
	particle int
	eid      int
	details
}

// Method 2
func (d details) promotmethod(tsalary int) int {
	return d.psalary * tsalary
}

type FinalSalary func(int, int) int

type Author2 struct {
	name      string
	language  string
	Marticles int
	Pay       int
	salary    FinalSalary
}

func main() {
	var a Address
	fmt.Println(a)

	a1 := Address{"Sudhir", "Ajmer", 305001}
	fmt.Println("Address1: ", a1)

	a2 := Address{Name: "Nisha", City: "Jaipur", Pincode: 304801}
	fmt.Println("Address2: ", a2)

	a3 := Address{Name: "Sunil"}
	fmt.Println("Address3: ", a3)

	a4 := Address{"Sudhir", "Ajmer", 305001}

	car1 := Car{Name: "Ferrari", Model: "GTC4", Color: "red", WeightInKg: 1920}
	fmt.Println("Car1 Name: ", car1.Name)
	fmt.Println("Car1 Color: ", car1.Color)

	car1.Color = "Black"
	fmt.Println("Car1 Color: ", car1)

	emp1 := &Employee{"Sam", "Anderson", 55, 60000}
	fmt.Println("Employee1 First Name: ", (*emp1).firstName)
	fmt.Println("Employee1 Age: ", (*emp1).age)

	// can also work like this
	fmt.Println("Employee1 First Name: ", emp1.firstName)
	fmt.Println("Employee1 Age: ", emp1.age)

	if a1 == a2 {
		fmt.Println("a1 is equal to a2")
	} else {
		fmt.Println("a1 is not equal to a2")
	}

	if a1 == a4 {
		fmt.Println("a1 is equal to a4")
	} else {
		fmt.Println("a1 is not equal to a4")
	}

	// Comparing a1 with a2
	// Using DeepEqual() method
	fmt.Println("Is a1 equal to a2: ", reflect.DeepEqual(a1, a2))

	// Comparing a1 with a4
	// Using DeepEqual() method
	fmt.Println("Is a1 equal to a4: ", reflect.DeepEqual(a1, a4))

	// Initializing the fields
	// of the structure
	// result := HR{
	// 	details: Author{"Sona", "ECE", 2013},
	// }

	// Display the values
	fmt.Println("\nDetails of Author")
	// fmt.Println(result)

	result := Teacher{
		name:    "Suman",
		subject: "Java",
		exp:     5,
		details: Student{"Bongo", "CSE", 2},
	}

	// Display the values
	fmt.Println("Details of the Teacher")
	fmt.Println("Teacher's name: ", result.name)
	fmt.Println("Subject: ", result.subject)
	fmt.Println("Experience: ", result.exp)

	fmt.Println("\nDetails of Student")
	fmt.Println("Student's name: ", result.details.name)
	fmt.Println("Student's branch name: ", result.details.branch)
	fmt.Println("Year: ", result.details.year)

	// Creating and initializing
	// the anonymous structure
	Element := struct {
		name      string
		branch    string
		language  string
		Particles int
	}{
		name:      "Pikachu",
		branch:    "ECE",
		language:  "C++",
		Particles: 498,
	}

	// Display the anonymous structure
	fmt.Println(Element)

	// Initializing the fields of
	// the student structure
	values := student{
		branch: "CSE",
		year:   2010,
		details: details{

			name:   "Sumit",
			age:    28,
			gender: "Male",
		},
	}

	// Promoted fields of the student structure
	fmt.Println("Name: ", values.name)
	fmt.Println("Age: ", values.age)
	fmt.Println("Gender: ", values.gender)

	// Normal fields of
	// the student structure
	fmt.Println("Year: ", values.year)
	fmt.Println("Branch : ", values.branch)

	// Initializing the fields of
	// the employee structure
	values2 := employee2{
		post:     "HR",
		eid:      4567,
		particle: 5,
		details: details{

			name:    "Sumit",
			age:     28,
			gender:  "Male",
			psalary: 890,
		},
	}

	// Promoted fields of
	// the employee structure
	fmt.Println("Name: ", values2.name)
	fmt.Println("Age: ", values2.age)
	fmt.Println("Gender: ", values2.gender)
	fmt.Println("Per day salary: ", values2.psalary)

	// Promoted method of
	// the employee structure
	fmt.Println("Total Salary: ", values2.promotmethod(30))

	// Normal fields of
	// the employee structure
	fmt.Println("Post: ", values2.post)
	fmt.Println("Employee id: ", values2.eid)
	fmt.Println("Total Articles: ", values2.promotmethod(30))

	// Initializing the fields
	// of the structure
	result3 := Author2{
		name:      "Sonia",
		language:  "Java",
		Marticles: 120,
		Pay:       500,
		salary: func(Ma int, pay int) int {
			return Ma * pay
		},
	}

	// Display values
	fmt.Println("Author's Name: ", result3.name)
	fmt.Println("Language: ", result3.language)
	fmt.Println("Total number of articles published in May: ", result3.Marticles)
	fmt.Println("Per article pay: ", result3.Pay)
	fmt.Println("Total salary: ", result3.salary(result3.Marticles, result3.Pay))
}
