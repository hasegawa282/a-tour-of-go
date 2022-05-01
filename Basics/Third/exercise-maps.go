package main

import (
	// "golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	result := make(map[string]int)
	arr := strings.Split(s, " ")
	for _, ar := range arr {
		result[ar] += 1
	}

	return result
}

func main() {
	// wc.Test(WordCount)
}
