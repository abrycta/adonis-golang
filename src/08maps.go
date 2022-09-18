package main

import "fmt"

func main() {
	//ages := make(map[string]int) // mapping from strings to ints
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	/*
		This is equivalent to
		ages := make(map[string]int)
		ages["alice"] = 31
		ages["charlie"] = 34
	*/

	//Map elements are accessed through the usual subscript notation:
	ages["alice"] = 32
	fmt.Println(ages["alice"]) // "32"

	//removed with the built-in function delete:
	delete(ages, "alice") // remove element ages["alice"]

	/*
		All of these operations are safe even if the element isn’t in the map; a map lookup using a key
		that isn’t present returns the zero value for its type, so, for instance, the following works even
		when "bob" is not yet a key in the map because the value of ages["bob"] will be 0.
	*/
	ages["bob"] = ages["bob"] + 1
	/*
		or
		ages["bob"] += 1
		or
		ages["bob"]++
	*/
	fmt.Println(ages["bob"])
	/* But a map element is not a variable, and we cannot take its address:

	_=&ages["bob"] // compile error: cannot take address of map element

	*/

}
