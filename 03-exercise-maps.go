package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	res := make(map[string]int, len(words))

	if len(s) == 0 {
		res["x"] = 1
	}

	for _, w := range words {
		res[w]++
	}

	return res
}

func main() {
	wc.Test(WordCount)
}
