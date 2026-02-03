package main

func findDiagonalOrder(nums [][]int) []int {
	res := []int{}
	queue := make([][2]int, 0)
	head := 0

	queue = append(queue, [2]int{0, 0})

	for head < len(queue) {
		curr := queue[head]
		head++

		r, c := curr[0], curr[1]
		res = append(res, nums[r][c])

		if c == 0 && r+1 < len(nums) {
			queue = append(queue, [2]int{r + 1, c})
		}

		if c+1 < len(nums[r]) {
			queue = append(queue, [2]int{r, c + 1})
		}
	}
	return res
}

func findDiagonalOrder2(nums [][]int) []int {
	groups := make(map[int][]int)

	for r := len(nums) - 1; r >= 0; r-- {
		for c := 0; c < len(nums[r]); c++ {
			diagonal := r + c
			groups[diagonal] = append(groups[diagonal], nums[r][c])
		}
	}

	maxKey := 0
	for k := range groups {
		if k > maxKey {
			maxKey = k
		}
	}

	res := []int{}
	for i := 0; i <= maxKey; i++ {
		res = append(res, groups[i]...)
	}
	return res
}
