package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}

	return lim
}

func main() {
	// Printlnは中のpowが2つ先に実行されたから呼び出される
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
