package main

import "fmt"

func vals() (int, int, string) { // this (int, int, string)
	return 1, 2, "Buckle my shoe" // the function returns 2 int and 1 string.

}

func main() {
	// Assign the return values of vals() to a, b, c
	a, b, c := vals()
	// output values
	fmt.Println("a:", a, "b:", b, "c:", c)
}
