package main

import "math"

func wallsAndGates(rooms [][]int) {
	rows, cols := len(rooms), len(rooms[0])
	queue := [][]int{}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if rooms[r][c] == 0 {
				queue = append(queue, []int{r, c})
			}
		}
	}

	dirs := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]
		for _, directions := range dirs {
			nr, nc := cell[0]+directions[0], cell[1]+directions[1]
			if nr < 0 || nr >= rows || nc < 0 || nc >= cols || rooms[nr][nc] == -1 || rooms[nr][nc] != int(math.Pow(2, 31))-1 {
				continue
			}
			rooms[nr][nc] = rooms[cell[0]][cell[1]] + 1
			queue = append(queue, []int{nr, nc})
		}
	}

}
