package main

import "fmt"

func findavg(a [3]int, size int) (r int) {
	var k, val int

	for k = 0; k < size; k++ {
		val += a[k]
	}

	r = val / size
	return
}

func main() {
	var myarr [3]string
	myarr[0] = "Sudhir"
	myarr[1] = "Meena"

	fmt.Println(myarr[0] + " " + myarr[1])

	arr := [4]string{"My", "Favourite", "is", "you"}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// multi dimensional array
	arr2 := [2][2]string{{"Python", "Golang"}, {"C++", "Java"}}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print(arr2[i][j] + " ")
		}
		fmt.Println()
	}

	arr3 := [3]int{9, 7, 6}
	arr4 := [...]int{9, 7, 6, 5, 4, 3, 2, 1}
	arr5 := [3]int{9, 3, 5}

	fmt.Println(len(arr3))
	fmt.Println(len(arr4))
	fmt.Println(len(arr5))

	// copy array by value, doesnt reflect changes to original array
	arr6 := arr5
	fmt.Println(arr5)
	fmt.Println(arr6)

	arr6[2] = 6
	fmt.Println(arr5)
	fmt.Println(arr6)

	// copy the array by reference, reflect changes to original array
	arr7 := &arr5
	fmt.Println(arr5)
	fmt.Println(*arr7)

	arr7[2] = 1000
	fmt.Println(arr5)
	fmt.Println(*arr7)

	fmt.Println(findavg(*arr7, 3))
}
