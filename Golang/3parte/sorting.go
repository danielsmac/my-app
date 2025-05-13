package main

import (
	"fmt"
	"slices"
)

func main() {

	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("Strings:", strs) // Output: [a b c]

	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Ints:", ints) // Output: [2 4 7]

	s := slices.IsSorted(ints)
	fmt.Println("Is sorted:", s) // Output: true
}