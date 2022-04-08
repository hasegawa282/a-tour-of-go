package main

import "fmt"

func main() {
	// return後に呼ばれる関数
	defer fmt.Println("world")

	fmt.Println("helllo")
}
