package main

import (
	"fmt"
	"sort"
)

// create a type that is an alias for the builtin []string type
type byLength []string

// implementation of sort.Interface - Len
func (s byLength) Len() int {
	return len(s)
}

// implementation of sort.Interface - Swap similar across type with Len
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// implementation of sort.Interface - Less, which holds the actual custom sorting logic
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	//implementation of custom sort by converting the names to byLength and then use sort.Sort on the typed names
	names := []string{"Clawrence", "Angelo", "Seifer"}
	sort.Sort(byLength(fruits))
	fmt.Println(names)
}
