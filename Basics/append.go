package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)

	b := [...]string{"Penn", "Teller"}
	printSliceS(b)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSliceS(s [2]string) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
