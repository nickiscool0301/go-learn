package main

func twoSum(nums []int, target int) []int {
	cache := make(map[int]int)
	for i, v := range nums {
		complement := target - v
		if j, exist := cache[complement]; exist {
			return []int{i, j}
		}
		cache[v] = i
	}
	return []int{-1, -1}
}

func main() {
	result := twoSum([]int{2, 7, 11, 15}, 9)
	println(result[0], result[1])
}
