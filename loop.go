package main

import "fmt"

func main() {
	i := 0
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}
	// Start at 0, Going up to 5, increment by 1
	for j := 0; j <= 5; j++ {
		fmt.Println(j)
	}
	var odd = "Print Odd Values only"
	fmt.Println(odd)
	// Only print out odd values
	for indx := 0; indx <= 10; indx++ {
		if indx%2 == 0 {
			continue
		}
		fmt.Println(indx)
	}
	for {
		fmt.Println("Keep Printing")
		break
	}
}
