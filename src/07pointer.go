package main

import "fmt"

// function that increments a variable that its argument points to
// and returns the pointer of the modified variable
// the * operator has 2 forms in the context of pointers:
//  1. in front of a base type, e.g.: *int
//  2. in front of a variable name, e.g.: *x
//
// For case 1, it means that something is expecting or returning
// a pointer to an int. Case 1 is a TYPE
// For case 2, you are accessing the variable of the pointer.
// Case 2 is an OPERATION
func incr(a *int) int {
	*a++ // DEREFERENCE a and increment the value stored
	// at that address in memory
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
