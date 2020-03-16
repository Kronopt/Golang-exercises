// Implement WordCount. It should return a map of the counts
// of each “word” in the string s. The wc.Test function runs
// a test suite against the provided function and prints
// success or failure.
//
// You might find strings.Fields helpful.
package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	word_counts := make(map[string]int)
	words := strings.Split(s, " ")

	for _, v := range words {
		_, ok := word_counts[v]
		if ok {
			word_counts[v] += 1
		} else {
			word_counts[v] = 1
		}
	}
	return word_counts
}

func main() {
	wc.Test(WordCount)
}
