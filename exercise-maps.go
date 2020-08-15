package main

import (
	"strings"
	
	"golang.org/x/tour/wc"
)

// WordCounts returns a map that counts the number of occurences of each word
// in the string s, considering that words in s are separated by one or many whitespaces
func WordCount(s string) map[string]int {
	m := map[string]int{}
	w := strings.Fields(s)
	for i := range w {
		v, ok := m[w[i]]
		if !ok {
			m[w[i]] = 1
		} else {
			m[w[i]] = v + 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
