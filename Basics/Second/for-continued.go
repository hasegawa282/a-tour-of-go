package main

import "fmt"

func main() {
	sum := 1
	// sumが2倍ずつ増える
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
