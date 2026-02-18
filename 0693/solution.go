package main

func hasAlternatingBits(n int) bool {
	prev := n & 1
	n >>= 1
	for n > 0 {
		curr := n & 1
		if prev == curr {
			return false
		}
		n >>= 1
		prev = curr
	}
	return true
}
