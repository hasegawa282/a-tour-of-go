package main

import (
	"fmt"
)

// func swap(x, y string) (string, string) {
// 	return y, x
// }

// func add(x int, y int) int {
// 	return x + y
// }

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// naked return
	return
}

func main() {
	fmt.Println(split(17))
}
