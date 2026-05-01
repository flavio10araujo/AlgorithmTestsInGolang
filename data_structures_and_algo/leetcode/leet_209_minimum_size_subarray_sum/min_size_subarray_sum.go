package leet_209_minimum_size_subarray_sum

func MinSubArrayLen(target int, nums []int) int {
	left := 0
	sum := 0
	size := 0

	for i, v := range nums {
		sum += v

		for {
			if sum < target {
				break
			}

			if size == 0 {
				size = (i - left) + 1
			} else {
				size = min(size, (i-left)+1)
			}

			sum -= nums[left]
			left++
		}
	}

	return size
}
