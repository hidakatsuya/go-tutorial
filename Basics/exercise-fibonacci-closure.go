package main

import "fmt"

func fibonacci() func() int {
	s := []int{}
	n := 0
	return func() int {
		var n_value int
		if n == 0 {
			n_value = 0
		} else if n == 1 {
			n_value = 1
		} else {
			n_value = s[n-2] + s[n-1]
		}
		s = append(s, n_value)
		n += 1
		return n_value
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
