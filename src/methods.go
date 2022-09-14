package main

import "fmt"

//method that has a receiver type of *rect
type rect struct {
    width, height int
}

func (r *rect) area() int {
    return r.width * r.height
}

func (r rect) perim() int {
    return 2*r.width + 2*r.height
}

func main() {
    r := rect{width: 10, height: 5}

    //call 2 methods defined for the struct
    fmt.Println("area: ", r.area())
    fmt.Println("perim:", r.perim())

    //use pointer receiver to avoid copying on method calls
    rp := &r
    fmt.Println("area: ", rp.area())
    fmt.Println("perim:", rp.perim())
}
