package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

func init() {
	fmt.Println("Welcome to init() function")
}

func init() {
	fmt.Println("Hello! init() function")
}

func main() {
	s := []int{345, 78, 123, 10, 76, 2, 567, 5}
	sort.Ints(s)
	fmt.Println("Sorted slice: ", s)

	// find index
	fmt.Println("Index value: ", strings.Index("Geeksforgeeks", "ks"))

	// find the time
	fmt.Println("Time: ", time.Now().Unix())
}
