package main

func canCompleteCircuit(gas []int, cost []int) int {
	totalGas, totalCost, diff := 0, 0, 0
	startIndex := 0
	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]
		diff += gas[i] - cost[i]
		if diff < 0 {
			diff = 0
			startIndex = i + 1
		}
	}
	if totalGas < totalCost {
		return -1
	}
	return startIndex
}
