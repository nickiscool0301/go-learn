package main

import (
	"math/rand"
)

type Solution struct {
	prefixSum []int
	totalSum  int
}

func Constructor(w []int) Solution {
	prefixSum := make([]int, len(w))
	totalSum := 0
	for i, num := range w {
		totalSum += num
		prefixSum[i] = totalSum
	}
	return Solution{
		prefixSum: prefixSum,
		totalSum:  totalSum,
	}
}

func (this *Solution) PickIndex() int {
	// rand.Intn: return random int in [0, n) -> have to plus 1
	// In python: self.total_sum * random.random(). random.random() return [0.0,1.0)

	target := rand.Intn(this.totalSum) + 1
	l, r := 0, len(this.prefixSum)-1
	res := 0
	for l <= r {
		m := l + (r-l)/2
		if this.prefixSum[m] >= target {
			res = m
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
