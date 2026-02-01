package main

func numPairsDivisibleBy60(time []int) int {
	res := 0
	modulos := make([]int, 60)
	for _, t := range time {
		modulo := t % 60
		complement := (60 - modulo) % 60
		res += modulos[complement]
		modulos[modulo]++
	}
	return res
}
