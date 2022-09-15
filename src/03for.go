package main

import "fmt"

func main() {
	// go only has one type of loop: the for loop
	i := 1

	// a traditional while loop
	for i <= 5 {
		fmt.Println(i)
		i = i + 1
	}

	// classic for loop
	// j := 6 is the initializer
	// j <= 9 is the condition
	// j++ is the increment
	for j := 6; j <= 9; j++ {
		fmt.Println(j)
	}

	// demonstrate control structure in
	// a for loop
	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
