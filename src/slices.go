package main

import (
	"fmt"
)

func main() {
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May",
		6: "June", 7: "July", 8: "August", 9: "September",
		10: "October", 11: "November", 12: "December"}
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)     // ["April" "May" "June"]
	fmt.Println(summer) // ["June" "July" "August"]
	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}
}

func array() {
	// var a1
	var a = [3]int{10, 20, 30} // array of 3 integers

	fmt.Println(a[0])        // print the first element
	fmt.Println(a[len(a)-1]) // print the last element, a[2]
	fmt.Println()
	// Print the indices and elements.
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	fmt.Println()
	// Print the elements only.
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
	fmt.Println()
	type Currency int

	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)
	symbol := [...]string{USD: "$", EUR: "9", GBP: "!", RMB: "¥"}
	fmt.Println(symbol[USD])
	//fmt.Println(RMB, symbol[RMB]) // "3 "¥"
}

func slice() {

}
