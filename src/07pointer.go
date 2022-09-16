package main

import "fmt"

// function that increments a variable that its argument points to and returns the pointer
// of the manipulated variable
func incr(a *int) int {
	*a++
	return *a
}

func main() {

	x := 1
	p := &x         //p, of type *int, points to x
	fmt.Println(*p) //prints "1"
	*p = 2          //assign/updates the value of x
	fmt.Println(x)  //prints the new value that is "2"

	a := 1
	incr(&a)              //side effect: a is now 2
	fmt.Println(incr(&a)) //prints the new value which is "3"
}
