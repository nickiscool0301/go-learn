package main

func totalFruit(fruits []int) int {
	freq := make(map[int]int)
	l, res := 0, 0
	for r := 0; r < len(fruits); r++ {
		freq[fruits[r]]++
		for len(freq) > 2 {
			freq[fruits[l]]--
			if freq[fruits[l]] == 0 {
				delete(freq, fruits[l])
			}
			l++
		}
		res = max(res, r-l+1)
	}
	return res
}
