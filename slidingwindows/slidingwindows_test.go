package slidingwindows

import "testing"

func TestMaxSubSum(t *testing.T) {
	tests := []struct {
		arr        []int
		windowSize int
		want       int
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

func TestSmallestSubArrayWithGivenSum(t *testing.T) {
	tests := []struct {
		arr  []int
		sum  int
		want int
	}{
		{arr: []int{3, 4, 1, 7, 6}, sum: 2, want: 0},
		{arr: []int{2, 1, 5, 2, 8}, sum: 8, want: 1},
		{arr: []int{2, 1, 5, 2, 3, 2}, sum: 7, want: 2},
		{arr: []int{3, 4, 1, 1, 6}, sum: 8, want: 3},
		{arr: []int{3}, sum: 3, want: 1},
		{arr: []int{3}, sum: 5, want: 0},
		{arr: []int{}, sum: 5, want: 0},
	}
	for _, tt := range tests {
		if got := SmallestSubArrayWithGivenSum(tt.arr, tt.sum); got != tt.want {
			t.Errorf("SmallestSubArrayWithGivenSum() = %v, want %v, arr: %v", got, tt.want, tt.arr)
		}
	}
}
