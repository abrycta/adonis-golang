package main

import (
    "fmt"
    "strings"
)

// Main function
func main() {

    //create a slice of strings
    slc := []string{"Clarence", "angelo", "lpn"}

    /*
    * Convert the case of the
    * elements in the slice of strings
    * using ToUpper() function
    */
    for x := 0; x < len(slc); x++ {

        //convert the case of the elements to upppercase
        res := strings.ToUpper(slc[x])

        // Exported Method
        fmt.Println(res)
    }
}
