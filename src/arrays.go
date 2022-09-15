package main

import "fmt"

// for range returns a copy of the index and element
// this guarantees immutability and reduces side effects
/*
func sumMatrix(arr *[2][3]int) {
	for rowIndex, row := range *arr {
		for col := range row {
			// := is for declaring NEW variables
			// this is invalid:
			// row[col] := rowIndex + ie
			row[col] = rowIndex + row[col]
		}
	}
}
*/

func main() {
	// initialize an array of size 5
	// valid indexes are 0 to (len - 1)
	// a[0] to a[4]
	var a [5]int // no uninitialized variables, array is initialized
	// to default values, that is all elements are filled with 0
	// because the default value of int is 0

	fmt.Println("show:", a)

	a[4] = 1 // assignment
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("length:", len(a))

	// array initializer expression
	b := [5]int{1, 2, 3, 4, 5} // assign b to an array literal
	fmt.Println("values:", b)

	// 2d-array demonstration
	twoD := [2][3]int{
		{2, 3, 4},
		{5, 6, 7},
	}

	fmt.Println("before: ", twoD)

	// add the rowIndex to the current element
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = twoD[i][j] + i
		}
	}
	fmt.Println("2d example: ", twoD)
}
