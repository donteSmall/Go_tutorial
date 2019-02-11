package main

import "fmt"

func main() {
	// Maps are similar to python dictionary (i.e. hashtable, etc)

	//specify key/value pair

	m := make(map[string]int)

	m["a"] = 0
	m["b"] = 1
	// Printing value for maps
	fmt.Println(m)

	//Printing the value for specific key
	fmt.Println(m["a"])

	//len() function is overloaded to work with maps:

	fmt.Println(len(m))

	//remove key/pair from map. (which can be done with delete keyword).
	delete(m, "a")
	fmt.Println(m)

	//another way to initialize a map (if you already know the values/keys)
	m2 := map[string]int{"key1": 1, "key2": 2}
	fmt.Println(m2)

	// Print the value and whether the value is present
	val, is_val_present := m["b"]
	fmt.Println(val)
	fmt.Println(is_val_present)
	// _ == means anything here, we really dont care
	_, is_val_present2 := m["a"]
	fmt.Println(is_val_present2)

}
