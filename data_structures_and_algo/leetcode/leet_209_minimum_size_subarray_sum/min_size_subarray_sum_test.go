package leet_209_minimum_size_subarray_sum

import (
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	got := MinSubArrayLen(11, []int{1, 2, 3, 4, 5})
	want := 3
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestMinSubArrayLenTableDriven(t *testing.T) {
	tests := []struct {
		target   int
		nums     []int
		expected int
	}{
		{7, []int{2, 3, 1, 2, 4, 3}, 2},
		{11, []int{1, 2, 3, 4, 5}, 3},
	}

	for _, tt := range tests {
		got := MinSubArrayLen(tt.target, tt.nums)
		if got != tt.expected {
			t.Errorf("got %d; want %d", got, tt.expected)
		}
	}
}
