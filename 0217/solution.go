package main

func containsDuplicate(nums []int) bool {
	uniques := make(map[int]bool)
	for _, num := range nums {
		if _, exist := uniques[num]; exist {
			return false
		}
		uniques[num] = true
	}
	return false
}
