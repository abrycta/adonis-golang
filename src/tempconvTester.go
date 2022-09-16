package main

import (
	"fmt"
	"src/tempconv" // Importing path of the package created
)

func main() { // Conversion tester
	fmt.Printf("Absolute Zero = %v\n", tempconv.AbsoluteZeroC) // Print the computed values Absolute Zero= -273.15°C"
	fmt.Println(tempconv.CToF(tempconv.BoilingC))              // "212°F"
	fmt.Println(tempconv.FToC(tempconv.FreezingC))             // "-17.77777777777778°C"
}
