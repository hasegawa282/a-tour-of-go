package main

import "fmt"

func main() {
	sum := 1
	// 「;」は省略できる
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
