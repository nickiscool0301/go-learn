package main

func binaryGap(n int) int {
	last := -1
	var res int
	for i := 0; i < 32; i++ {
		if (n>>i)&1 == 1 {
			if last != -1 {
				res = max(res, i-last)
			}
			last = i
		}
	}
	return res

}
