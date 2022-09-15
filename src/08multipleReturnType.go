package main

import "fmt"

func vals() (int, int, string) { // this (int, int, string) in this function indicates that the function returns 2 int and 1 string.
	return 1, 2, "Buckle my shoe"
}

func main() {

	a, b, c := vals() // Variable declaration of a, b, and c is equivalent to vals() function.

	/*
		prints variables a, b, and c
	*/
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
