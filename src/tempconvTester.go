package main

import (
	"fmt"
	"src/tempconv"
)

func main() {
	fmt.Printf("Absolute Zero = %v\n", tempconv.AbsoluteZeroC) // "(Conversion) Absolute Zero= -273.15°C"
	fmt.Println(tempconv.CToF(tempconv.BoilingC))              // (Conversion) "212°F"
	fmt.Println(tempconv.FToC(tempconv.FreezingC))             // (Conversion) "-17.77777777777778°C"
}
