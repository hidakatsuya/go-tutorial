package main

import "fmt"

func main() {
	var sum = 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println(sum)

	var sums = fmt.Sprint(sum) + "str"
	fmt.Println(sums)
}
