package main

import "fmt"

// function that takes two ints and return their product
func product(a int, b int) int {
	return a * b
}

// function that takes two ints and return their sum
func sum(a int, b int) int {
	return a + b
}

// function that takes three ints and return their sum
func sumOfThree(a int, b int, c int) int {
	return a + b + c
}

// TODO ADD POINTER FUNCTIONS
func main() {

	//call a function using name(args)
	resultSum := sum(5, 10)
	fmt.Println("5+10 =", resultSum)

	resultProduct := product(4, 5)
	fmt.Println("4*5 =", resultProduct)

	resultSumOfThree := sumOfThree(5, 10, 15)
	fmt.Println("5+10+15 =", resultSumOfThree)

}
