package main

func canSeePersonsCount(heights []int) []int {
	n := len(heights)
	res := make([]int, n)
	stack := []int{}

	for i := n - 1; i >= 0; i-- {
		for len(stack) != 0 && heights[stack[len(stack)-1]] < heights[i] {
			res[i]++
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			res[i]++
		}
		stack = append(stack, i)
	}
	return res
}
