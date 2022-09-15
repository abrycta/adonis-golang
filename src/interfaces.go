package main

import (
	"fmt"
	"math"
)

/*
interface for geometric shapes
*/
type geometry interface {
	area() float64
	perim() float64
}

/*
implementing the interface on geometric shapes circle and rectangle
*/
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

/*
Implementation for rectangle
*/
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

/*
Implementation for circle
*/
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

/*
calling methods in named interface
*/
func measure(g geometry) {
	fmt.Println(g)         // prints the inserted values
	fmt.Println(g.area())  // prints the area of the geometric shape
	fmt.Println(g.perim()) // prints the perimeter of the geometric shape
}

/*
Main method where cirle and rectangle implements the geometry interface as arguments to function measure.
*/
func main() {
	r := rect{width: 3, height: 4} // input height and width of a rectangle
	c := circle{radius: 5}         // input radius of a circle

	measure(r) // call measure function
	measure(c)
}
