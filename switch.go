package main

import "fmt"

func main() {
	trigger := 3
	switch trigger {
	case 1:
		fmt.Println("Case one")
	case 2:
		fmt.Println("Case two")
	case 3:
		fmt.Println("Case three")
	}

	trigger_Val := 5
	switch trigger_Val {
	case 1, 2:
		fmt.Println("one or two")
	case 3:
		fmt.Println("Case three")
	default:
		fmt.Println("Not one, two, or three")
	}

	switch {
	case trigger_Val == 5:
		fmt.Println(trigger_Val, "is equal to 5")
	case trigger_Val < 5:
		fmt.Println(trigger_Val, "Is less than 5")
	case trigger_Val > 5:
		fmt.Println(trigger_Val, "Is greater than 5")
	}

}
