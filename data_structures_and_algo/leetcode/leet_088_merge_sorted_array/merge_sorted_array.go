package leet_088_merge_sorted_array

func Merge(nums1 []int, m int, nums2 []int, n int) {
	left := m - 1
	right := n - 1
	next := m + n - 1

	for left >= 0 && right >= 0 {
		if nums1[left] >= nums2[right] {
			nums1[next] = nums1[left]
			left--
		} else {
			nums1[next] = nums2[right]
			right--
		}
		next--
	}

	for right >= 0 {
		nums1[next] = nums2[right]
		next--
		right--
	}
}
