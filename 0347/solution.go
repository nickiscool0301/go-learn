package main

func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	n := len(nums)

	for _, num := range nums {
		count[num]++
	}

	buckets := make([][]int, n+1)
	for num, freq := range count {
		buckets[freq] = append(buckets[freq], num)
	}

	res := make([]int, 0, k)

	for i := n; i >= 0; i-- {
		for _, num := range buckets[i] {
			res = append(res, num)
			if len(res) == k {
				return res
			}
		}
	}
	panic("should never reach here")
}
