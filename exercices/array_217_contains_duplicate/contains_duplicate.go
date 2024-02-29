package main

/*
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.
Example 1: Input: nums = [1,2,3,1] Output: true
Example 2: Input: nums = [1,2,3,4] Output: false
Example 3: Input: nums = [1,1,1,3,3,4,3,2,4,2] Output: true
*/
func main() {
	//var nums = []int{1, 2, 3, 1}
	//var nums = []int{1, 2, 3, 4}
	//var nums = []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	var nums = []int{1, 2, 2, 4}
	var ret = check(nums)
	print(ret)
}

func check(nums []int) bool {
	// solution 01: using a set to keep the elements and check if the element is already in the set, if yes return false.
	// T: O(n) S: O(n)

	// solution 02: sort the array and iterate
	// T: O(n log n) S: O(1)

	var myset = map[int]bool{}

	for i := range nums {
		_, ok := myset[nums[i]]
		if ok {
			return true
		}

		myset[nums[i]] = true
	}

	return false
}
