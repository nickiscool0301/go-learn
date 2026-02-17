package main

import (
	"fmt"
	"math/bits"
)

func readBinaryWatch(turnedOn int) []string {
	var res []string
	for i := 0; i < 1024; i++ {
		h := i >> 6
		m := i & 0x3F
		if h < 12 && m < 60 && bits.OnesCount(uint(i)) == turnedOn {
			res = append(res, fmt.Sprintf("%d:%02d", h, m))
		}
	}
	return res
}
