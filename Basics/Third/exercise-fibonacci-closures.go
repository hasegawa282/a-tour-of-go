package main

import "fmt"

func fibonacci() func() int {
	nums := []int{}
	return func() int {
		len_nums := len(nums)
		if len_nums == 0 {
			nums = append(nums, 0)
		} else if len_nums == 1 {
			nums = append(nums, 1)
		} else {
			nums = append(nums, nums[len_nums-2]+nums[len_nums-1])
		}
		return nums[len(nums)-1]
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
