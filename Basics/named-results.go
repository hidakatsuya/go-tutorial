package main

import "fmt"

func split(sum int) (a, b int) {
	a = sum * 4 / 9
	b = sum - a
	return
}

func main() {
	fmt.Println(split(17))
}
