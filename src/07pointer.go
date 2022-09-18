package main

import "fmt"

func main() {

	x := 1
	var p *int               //*int = base type
	p = &x                   //p, of type *int, points to x
	fmt.Println("*p = ", *p) //prints "1"
	*p = 2                   //assign/updates the value of x
	fmt.Println("x = ", x)   //prints the new value that is "2"
}
