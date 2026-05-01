package leet_088_merge_sorted_array

import (
	"slices"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		nums1    []int
		m        int
		nums2    []int
		n        int
		expected []int
	}{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
	}

	for _, tt := range tests {
		Merge(tt.nums1, tt.m, tt.nums2, tt.n)
		if !slices.Equal(tt.nums1, tt.expected) {
			t.Errorf("got %v, want %v", tt.nums1, tt.expected)
		}
	}
}
