package main

import "fmt"

func moveZeroes(nums []int) {
	insertPos := 0
	for i, num := range nums {
		if num != 0 {
			nums[i], nums[insertPos] = nums[insertPos], nums[i]
			insertPos++
		}
	}
}

func main() {
	nums := []int{0, 2, 0, 1, 2, 0, 10, 15}
	moveZeroes(nums)
	fmt.Print(nums)
}
