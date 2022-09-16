// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv // creation and naming of a package

import "fmt"

type Celsius float64 //variable declarations
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15 // Initialization of constant variables
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) } // Functions that return the computed values to string
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) } // Functions that compute the conversion of the following

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Celsius) Celsius { return Celsius((f - 32) * 5 / 9) }
