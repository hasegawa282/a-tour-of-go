package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		// 後に入れたものを先に出す
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
