package main

import "fmt"

func main() {
	s := make([]int, 3)

	//Are similar to arrays
	s[0] = 1
	s[1] = 2
	s[2] = 3

	//Print the entire array
	fmt.Println(s)
	//print the value at the 0 indx
	fmt.Println(s[0])

	//print the length of the array
	fmt.Println(len(s))

	// Append function is unique to slices
	s = append(s, 5)
	fmt.Println(s)
	// Appends mult values to array
	s = append(s, 5, 6, 7, 8)
	fmt.Println(s)

	//slice syntax
	fmt.Println(s[1:3])

	fmt.Println(s[:2])

	// concise slice defintion
	t := []int{100, 200, 300}
	fmt.Println(t)

	x := s
	fmt.Println(x)

	x[0] = 500

	// Notice the first value in both array changed.
	fmt.Println(x)
	fmt.Println(s)

}
