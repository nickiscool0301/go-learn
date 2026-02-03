package main

/*
Idea: 2 turning points -> opposite sign -> use multiply to check
*/

func isTrionic(nums []int) bool {
	n := len(nums)
	if nums[0] >= nums[1] {
		return false
	}
	count := 1
	for i := 2; i < n; i++ {
		if nums[i-1] == nums[i] {
			return false
		}
		if (nums[i-1]-nums[i-2])*(nums[i]-nums[i-1]) < 0 {
			count++
		}
	}
	return count == 3
}
