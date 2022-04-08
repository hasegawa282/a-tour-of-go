package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 1; i <= 10000; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func main() {
	// 自作平方根
	sq := Sqrt(2)
	// mathライブラリ平方根
	sq_math := math.Sqrt(2)
	fmt.Println(sq, sq_math)
}
