package main

import "math"

func minimumCardPickup(cards []int) int {
	seen := make(map[int]int, len(cards))
	res := math.MaxInt
	for i, value := range cards {
		if prevCard, exist := seen[value]; exist {
			if i-prevCard+1 < res {
				res = i - prevCard + 1
			}
		}
		seen[value] = i
	}
	if res == math.MaxInt {
		return -1
	}
	return res
}
