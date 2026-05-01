package main

import (
	"sort"
)

var maxweight = 7
var max = 0

func main() {
	var weightSum = 0
	values := [][]int{{3, 4}, {4, 5}, {7, 8}}
	//values := [][]int{{1, 100}, {4, 5}, {6, 8}, {5, 90}}
	//print(dp(max, values))

	for i := 0; i < len(values); i++ {
		dfs(values, i, weightSum, 0)
	}

	print(max)
}

func dfs(values [][]int, n int, weightSum int, maxSum int) int {
	if n >= len(values) {
		return 0
	}

	if weightSum+values[n][0] > maxweight {
		return 0
	}

	weightSum = weightSum + values[n][0]
	maxSum = maxSum + values[n][1]

	if maxSum > max {
		max = maxSum
	}

	for i := n + 1; i < len(values); i++ {
		dfs(values, i, weightSum, maxSum)
	}

	return 0
}

func dp(max int, values [][]int) int {
	sort.Slice(values, func(i, j int) bool {
		return values[i][1] > values[j][1]
	})

	for i, value := range values {
		println(i, value[1])
	}

	return 0
}
