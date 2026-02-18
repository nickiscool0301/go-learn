package main

func numIslands(grid [][]byte) int {
	rows, cols := len(grid), len(grid[0])

	count := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				dfs(grid, r, c)
				count++
			}
		}
	}
	return count
}

func dfs(grid [][]byte, row int, col int) {
	rows, cols := len(grid), len(grid[0])
	if row < 0 || row >= rows || col < 0 || col >= cols || grid[row][col] == '0' {
		return
	}
	grid[row][col] = '0'
	dirs := [4][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	for _, d := range dirs {
		dfs(grid, row+d[0], col+d[1])
	}
}
