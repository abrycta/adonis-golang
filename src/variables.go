package main

import "fmt"

var l int = 24 // full variable declaration syntax at package level

func main() {

	/*
		different methods to instantiate a variable
	*/
	var i int
	i = 1
	var j int = 1 // full variable declaration syntax
	k := 2        //

	fmt.Printf("%v, %T\n", i, i) //prints value of variable and variable type
	fmt.Println(j, k)            //variables must always be used in Go, otherwise you get a runtime error.
}
