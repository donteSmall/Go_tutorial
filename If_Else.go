package main

import "fmt"

func main() {
	indx := 5
	if indx%2 == 0 {
		fmt.Println(indx, "is even")

	} else {
		fmt.Println(indx, "is even")
	}

	num := 100
	if num == 100 {
		fmt.Println(num, "is 100")

	}
	value := 55
	if value < 50 {
		fmt.Println(value, "is less than 50")
	} else if value > 50 {
		fmt.Println(value, "is greater than 50")
	} else {
		fmt.Println(value, "is equal to 50 ")
	}
}
