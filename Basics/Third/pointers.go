package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	*p = *p / 37
	fmt.Println(j)
}
