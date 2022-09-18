package main

/*
// UNCOMMENT THIS CODE OUT IF DEMONSTRATING

import (
	"errors"
	"fmt"
	"os"
)

// ERROR HANDLING WITH PANIC AND RECOVER
func catchPanic() {
	// recover returns an errorval if
	// a call to panic() was raised
	// otherwise, recover() returns nil
	err := recover()

	// test for the existence of an error
	if err != nil {
		fmt.Println(err, "recovered.")
		fmt.Println("exiting program normally")
		os.Exit(0)
	}
}

func quotientMod(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		panic("Denominator is 0")
	}
	return numerator / denominator, numerator % denominator, nil
}

func main() {
	numerator := 20
	denominator := 0
	quotient, mod, err := quotientMod(numerator, denominator)

	// err is now declared as a variable
	// it will be a compile-time error to not utilize the variable
	// go forces you to handle errors gracefully
	if err != nil { // block is executed when denominator is 0
		fmt.Println(err) // outputs "denominator is 0" to the console
		os.Exit(1)
	}
	fmt.Println("quotient:", quotient, "mod:", mod)
}

*/
