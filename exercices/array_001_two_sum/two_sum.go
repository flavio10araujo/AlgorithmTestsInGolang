package main

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.

Example 1:
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:
Input: nums = [3,2,4], target = 6
Output: [1,2]

Example 3:
Input: nums = [3,3], target = 6
Output: [0,1]
*/

func main() {
	//var target = 9
	//var nums = []int{2, 7, 11, 15}

	//var target = 6
	//var nums = []int{3, 2, 4}

	var target = 6
	var nums = []int{3, 3}

	var ret = make([]int, 2)

	// Solution 01: two fors, brute force. T=O(n ^ 2); S=O(1)

	// Solution 02: sort the array, use two pointers opposite ways. T=O(n log n); S=O(1)

	// Solution 03: iterate the array and add each number and its index in a map[int][]int; iterate again and look for (target - nums[i]) in the map; then check the index
	// T=O(n); S=O(n)

	// Solution 03b: iterate the array adding in the map and checking if the complement is in the map, everything in the same for.

	ret = twoSum(nums, target)

	print(ret[0], " : ", ret[1])
}

func twoSum(nums []int, target int) []int {
	var mymap = make(map[int]int)

	for i, num := range nums {
		if index, ok := mymap[target-num]; ok {
			return []int{index, i}
		}

		mymap[num] = i
	}

	return nil
}
