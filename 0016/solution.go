package main

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}
	closest := nums[0] + nums[1] + nums[2]
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, n-1
		for l < r {
			currSum := nums[i] + nums[l] + nums[r]
			if currSum == target {
				return currSum
			}
			if abs(currSum-target) < abs(closest-target) {
				closest = currSum
			} else if currSum > target {
				r--
			} else {
				l++
			}
		}
	}
	return closest
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
