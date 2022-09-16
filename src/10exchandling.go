package main

import (
	"errors"
	"fmt"
	"os"
)

func errorfunc(arg int) (int, error) { // for error declaration
	if arg == 10 {
		return -1, errors.New("Error occurred.")
	}
	return arg, nil
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

	for _, i := range []int{4, 10} { // error function
		if r, e := errorfunc(i); e != nil {
			fmt.Println("function failed:", e) // returns an error
		} else {
			fmt.Println("function worked:", r) // returns a value under 10
		}
	}

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
