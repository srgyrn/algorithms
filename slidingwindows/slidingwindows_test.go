package slidingwindows

import "testing"

func TestMaxSubSum(t *testing.T) {
	tests := []struct {
		arr  []int
		windowSize int
		want int
	}{
		{arr: []int{2, 1, 5, 1, 3, 2}, windowSize: 3, want: 9},
		{arr: []int{2, 3, 4, 1, 5}, windowSize: 2, want: 7},
	}
	for _, tt := range tests {
		if got := MaxSubSum(tt.arr, tt.windowSize); got != tt.want {
			t.Errorf("MaxSubSum(%v, %v) = %v, want %v", tt.arr, tt.windowSize, got, tt.want)
		}
	}
}
