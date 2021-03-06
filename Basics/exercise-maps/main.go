package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	count := make(map[string]int)

	for _, word := range words {
		// count[word] = count[word] + 1
		count[word]++
	}
	return count
}

func main() {
	wc.Test(WordCount)
}
