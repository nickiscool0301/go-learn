package main

func reverseBits(n int) int {
	var res int
	for i := 0; i < 32; i++ {
		res <<= 1
		res = res | (n & 1)
		n >>= 1
	}
	return res
}
