package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	var a []string = names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)

	names[0] = "ZZZ"
	names[2] = "YYY"
	fmt.Println(a, b)
	fmt.Println(names)
}
