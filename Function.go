package main

import "fmt"

//  (specify parameters) and return type
func add(a int, b int) int {
	return a + b
}

// if more than one parameters, can specify at the end
func add3(a, b, c int) int {
	return a + b + c
}
func main() {
	ans := add(1, 1)
	fmt.Println("Answer -> ", ans)

	ans2 := add3(1, 1, 3)
	fmt.Println("Answer -> ", ans2)

}
