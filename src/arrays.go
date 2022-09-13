package main

import "fmt"

func main() {
    var a [5]int // initialize array
    fmt.Println("show:", a)
    a[4] = 1
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])
    fmt.Println("length:", len(a)) 
    b := [5]int{1, 2, 3, 4, 5} //to create values
    fmt.Println("values:", b)
  
    var twoD [2][3]int // to create 2d array
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d example: ", twoD)
}
