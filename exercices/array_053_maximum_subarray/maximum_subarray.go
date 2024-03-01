package main

/*
Given an integer array nums, find the subarray with the largest sum, and return its sum.

Example 1:
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: The subarray [4,-1,2,1] has the largest sum 6.

Example 2:
Input: nums = [1]
Output: 1
Explanation: The subarray [1] has the largest sum 1.

Example 3:
Input: nums = [5,4,-1,7,8]
Output: 23
Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.
*/
func main() {
	var nums = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	//var nums = []int{1}
	//var nums = []int{5, 4, -1, 7, 8}
	print(maxSubarray(nums))
}

func maxSubarray(nums []int) int {

	// Solution 01: two fors, brute force. T: O(n ^ 2). S: O(1).

	// Solution 02:

	var maxTotal = nums[0]
	var maxWindow = nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i]+maxWindow {
			maxWindow = nums[i]
		} else {
			maxWindow = maxWindow + nums[i]
		}

		if maxWindow > maxTotal {
			maxTotal = maxWindow
		}
	}

	return maxTotal
}