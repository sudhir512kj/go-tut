package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := [7]string{"This", "is", "the", "tutorial",
		"of", "Go", "language"}
	// Display array
	fmt.Println("Array:", arr)

	// Creating a slice
	myslice := arr[1:6]

	// Display slice
	fmt.Println("Slice:", myslice)

	// Display length of the slice
	fmt.Printf("Length of the slice: %d", len(myslice)) // 5

	// Display the capacity of the slice
	fmt.Printf("\nCapacity of the slice: %d\n", cap(myslice)) // 6

	// Creating a slice
	// using the var keyword
	var my_slice_1 = []string{"Geeks", "for", "Geeks"}

	fmt.Println("My Slice 1:", my_slice_1)

	// Creating a slice
	//using shorthand declaration
	my_slice_2 := []int{12, 45, 67, 56, 43, 34, 45}
	fmt.Println("My Slice 2:", my_slice_2)
	fmt.Println(cap(my_slice_2))
	my_slice_2 = append(my_slice_2, 20)
	fmt.Println("My Slice 2:", my_slice_2)

	var my_slice_3 = make([]int, 4, 7)
	fmt.Printf("Slice 3 = %v, \nlength = %d, \ncapacity = %d\n",
		my_slice_3, len(my_slice_3), cap(my_slice_3))

	// Creating another array of size 7
	// and return the reference of the slice
	// Using make function
	var my_slice_4 = make([]int, 7)
	fmt.Printf("Slice 4 = %v, \nlength = %d, \ncapacity = %d\n",
		my_slice_4, len(my_slice_4), cap(my_slice_4))

	for index, ele := range my_slice_2 {
		fmt.Println(ele, index)
	}

	// Creating multi-dimensional slice
	s2 := [][]string{
		{"Geeks", "for"},
		{"Geeks", "GFG"},
		{"gfg", "geek"},
	}

	// Accessing multi-dimensional slice
	fmt.Println("Slice 2 : ", s2)

	// Performing sort operation on the
	// slice using sort function
	sort.Strings(myslice)
	sort.Ints(my_slice_2)

	fmt.Println("\nAfter sorting:")
	fmt.Println("Slice 1: ", myslice)
	fmt.Println("Slice 2: ", my_slice_2)

	fmt.Println("\n=====Copy Slices=====\n")
	slc1 := []int{58, 69, 40, 45, 11, 56, 67, 21, 65}
	var slc2 []int
	slc3 := make([]int, 5)
	slc4 := []int{78, 50, 67, 77}

	// Before copying
	fmt.Println("Slice_1:", slc1)
	fmt.Println("Slice_2:", slc2)
	fmt.Println("Slice_3:", slc3)
	fmt.Println("Slice_4:", slc4)

	// Copying the slices
	copy_1 := copy(slc2, slc1)
	fmt.Println("\nSlice:", slc2)
	fmt.Println("Total number of elements copied:", copy_1)

	copy_2 := copy(slc3, slc1)
	fmt.Println("\nSlice:", slc3)
	fmt.Println("Total number of elements copied:", copy_2)

	copy_3 := copy(slc4, slc1)
	fmt.Println("\nSlice:", slc4)
	fmt.Println("Total number of elements copied:", copy_3)

	// Don't confuse here, because in above
	// line of code the slc4 has been copied
	// and hence modified permanently i.e.
	// slc 4 contains [58 69 40 45]
	copy_4 := copy(slc1, slc4)
	fmt.Println("\nSlice:", slc1)
	fmt.Println("Total number of elements copied:", copy_4)
}
