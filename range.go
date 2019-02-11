package main

import "fmt"

func main() {
	//range is similar to range in Python
	str_arr := []string{"a", "b", "c", "d"}
	for idx, char := range str_arr {
		fmt.Println("idx =", idx)
		fmt.Println("Char ->", char)
	}

	// Can range over key/value pairs in map
	a_map := map[string]int{"a": 1, "b": 2}
	for k, val := range a_map {
		fmt.Println("Key ->", k, "Val = ", val)
	}
	//can also just range over the keys in maps:
	for k := range a_map {
		fmt.Println("key:", k)
	}
}
