package main

/*
You are given an array prices where prices[i] is the price of a given stock on the ith day.
You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

Example 1:
Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.

Example 2:
Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.
*/
func main() {
	//var prices = []int{100}
	//var prices = []int{7, 1, 5, 3, 6, 4}
	//var prices = []int{7, 6, 5, 4, 3, 1}
	var prices = []int{1, 9, 0, 9}
	print(getRich(prices))
}

func getRich(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	var maxProfit = 0
	var L = 0
	var R = 0

	// Solution 01: two fors. T: O(n ^ 2); S: O(1).

	// Solution 02: two pointers same direction. T: O(n); S: O(1).

	for R < len(prices) {
		if prices[L] < prices[R] {
			if (prices[R] - prices[L]) > maxProfit {
				maxProfit = prices[R] - prices[L]
			}

			R++
		} else {
			L = R
			R++
		}
	}

	return maxProfit
}
