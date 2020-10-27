package main

import "fmt"

func findDup(nums []int) int {
	defer fmt.Print("Your duplicate value is: ")
	tort := nums[0]
	hare := nums[0]
	for {
		tort = nums[tort]
		hare = nums[nums[hare]]
		if tort == hare {
			break
		}
	}

	r1 := nums[0]
	r2 := tort
	for r1 != r2 {
		r1 = nums[r1]
		r2 = nums[r2]
	}
	return r1
}

func main() {
	test := []int{6, 2, 9, 4, 3, 1, 7, 1, 8, 5}
	fmt.Print(findDup(test))
}
