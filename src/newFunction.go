package main

import (
	"errors"
	"fmt"
	"os"
)

func declare() { // for declaring a function
	fmt.Println("Simple declaration")
}

func plus(a int, b int) int { // for function values addition
	return a + b
}

func minus(a int, b int) int { // for function values multiplication
	return a * b
}

func recur(n int) int { // for recursion
	if n == 0 {
		return 1
	}
	return n * recur(n-1)
}

func multivalue() (int, int) { // for multiple return value
	return 1, 2
}

func errorfunc(arg int) (int, error) { // for error declaration
	if arg == 10 {
		return -1, errors.New("Error occurred.")
	}
	return arg, nil
}

func variadic(nums ...int) { // for variadic functions
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func create(p string) *os.File { // creating file (defer)
	fmt.Println("Creating (deferred function) ")
	f, _ := os.Create(p)
	return f
}

func close(f *os.File) { // closing file (defer)
	fmt.Println("Closing (deferred function) ")
}

func panicfunc() { // panic function
	panic("a problem")
}

func main() {
	declare() // declare a function

	res := plus(1, 2)
	fmt.Println("addition: ", res) // simple function of adding values

	res = minus(2, 1)
	fmt.Println("multiplication: ", res) // simple function of multiplying values

	fmt.Println("simple recursion: ", recur(7)) // function of recursion

	a, b := multivalue()
	fmt.Print(a, b, " ")
	_, c := multivalue()
	fmt.Println(c) // function of multiple return values with blank identifier

	for _, i := range []int{4, 10} { // error function
		if r, e := errorfunc(i); e != nil {
			fmt.Println("function failed:", e) // returns an error
		} else {
			fmt.Println("function worked:", r) // returns a value under 10
		}
	}

	variadic(1, 2, 3, 4) // variadic function of sums

	f := create("tmp/defer.txt") // to call a deferred function
	defer close(f)               // to execute a close deferred function

	_, prob := os.Create("/tmp/file") // creating an error
	if prob != nil {
		panic(prob)
	}
	panic("problem occurred") // panic function

	defer func() { // recover function
		if r := recover(); r != nil {
			fmt.Println("Recovered error:\n ", r)
		}
	}()
	panicfunc()
}
