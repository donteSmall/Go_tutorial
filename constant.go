package main

import "fmt"

func main() {
	const pi float64 = 3.1415926

	const c int = 1000

	fmt.Println(c)
	// cannot assign to constant value
	// c = 5

	var d int = 100

	fmt.Println(d)

	d = 5000

	fmt.Println(d)

}
