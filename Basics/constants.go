package main

import "fmt"

const Pi = 3.14

func main() {
	const w = "世界"
	fmt.Println("Hello", w)
	fmt.Println("Happy", Pi)
	const Truth = true
	fmt.Println("Go rules?", Truth)
}
