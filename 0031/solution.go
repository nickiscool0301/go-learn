package main

func nextPermutation(nums []int) {
	n := len(nums)
	lastIndex := n - 2
	// find turning point
	for lastIndex >= 0 && nums[lastIndex] >= nums[lastIndex+1] {
		lastIndex--
	}
	if lastIndex >= 0 {
		pivot := n - 1
		for pivot > 0 && nums[pivot] <= nums[lastIndex] {
			pivot--
		}
		nums[lastIndex], nums[pivot] = nums[pivot], nums[lastIndex]
	}
	lastIndex = lastIndex + 1
	swapIndex := n - 1
	for lastIndex < swapIndex {
		nums[lastIndex], nums[swapIndex] = nums[swapIndex], nums[lastIndex]
		lastIndex++
		swapIndex--
	}
}
