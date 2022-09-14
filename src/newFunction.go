package main

import (
	"fmt"
	"math"
)

func hypot(x, y float64) float64 { 
	return math.Sqrt(x*x + y*y)
}

func fact(n int) int{ 
  if n == 0 {
       return 1
    }
  return n * fact(n-1)
}

func vals() (int, int) {
    return 3, 7
}

func main() {

// fuction by declaration
	fmt.Print("float value: ")
	fmt.Println(hypot(3, 4))
  
// fuction by recursion
  var fib func(n int) int
  fib = func(n int) int {
    if n < 2 {
       return n
    }
    return fib(n-1) + fib(n-2)
  }
  fmt.Print("fib value: ")
  fmt.Println(fib(7))
  
// fuction by multiple return values
  fmt.Print("multiple return value: ")
  a, b := vals()
  fmt.Print(a)
  fmt.Print(", ")
  fmt.Print(b)
  fmt.Print(", ")
  _, c := vals()
  fmt.Print(", ")
  fmt.Print(c)
  fmt.Println()
  
// functions by error handling
  
// functions by anonymous values
  
// functions by variadic functions
  
// functions by deferred function calls
  
// functions by panic
  
// functions by recover
  
}
