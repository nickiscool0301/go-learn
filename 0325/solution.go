package main

func maxSubArrayLen(nums []int, k int) int {
	seen := make(map[int]int, len(nums))
	seen[0] = -1

	currSum, res := 0, 0
	for i, num := range nums {
		currSum += num

		if prev, ok := seen[currSum-k]; ok {
			res = max(res, i-prev)
		}

		if _, ok := seen[currSum]; !ok {
			seen[currSum] = i
		}
	}
	return res
}
