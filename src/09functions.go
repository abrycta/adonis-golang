package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func multiply(a int, b int) int {
	return a * b
}

func recur(n int) int {
	if n == 0 {
		return 1
	}
	return n * recur(n-1)
}

func multivalue() (int, int) {
	return 1, 2
}

// nums is declared as an
func variadicSums(nums ...int) { // for variadic functions
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	res := add(1, 2)
	fmt.Println("addition:", res)

	// res := multiply is invalid
	// res already declared in scope
	// use assignment notation =
	res = multiply(2, 1)
	fmt.Println("multiplication:", res)

	fmt.Println("simple recursion:", recur(7))

	a, b := multivalue() // returns 1 & 2, respectively
	fmt.Println("a:", a, "b:", b)
	_, c := multivalue() // multiple return values with blank identifier
	fmt.Println(c)       // c = 2 , 1 is discarded

	variadicSums(1, 2, 3, 4) // can be of an arbitrary length
}
