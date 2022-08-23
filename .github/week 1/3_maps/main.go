package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := make(map[string]int)
	for _, word := range strings.Fields(s) {
		words[word]++
	}
	return words
}

func main() {
	wc.Test(WordCount)
}
