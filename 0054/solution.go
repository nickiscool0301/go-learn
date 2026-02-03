package main

func spiralOrder(matrix [][]int) []int {
	ROWS, COLS := len(matrix), len(matrix[0])
	top, bottom, left, right := 0, ROWS-1, 0, COLS-1
	res := []int{}

	for top <= bottom && left <= right {
		for i := left; i < right+1; i++ {
			res = append(res, matrix[top][i])
		}
		top++

		for i := top; i < bottom+1; i++ {
			res = append(res, matrix[i][right])
		}
		right--

		if top <= bottom {
			for i := right; i > left-1; i-- {
				res = append(res, matrix[bottom][i])
			}
			bottom--
		}

		if left <= right {
			for i := bottom; i > top-1; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}
	}
	return res
}
