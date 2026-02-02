package main

func searchMatrix(matrix [][]int, target int) bool {
	ROWS, COLS := len(matrix), len(matrix[0])
	l, r := 0, ROWS*COLS-1
	for l <= r {
		m := (l + r) / 2
		row := m / COLS
		col := m % COLS
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return false
}
