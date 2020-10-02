package slidingwindows

// MaxSubSum finds the maximum sum of any contiguous subarray of size "windowSize" in the int slice given
func MaxSubSum(arr []int, windowSize int) int {
	windowSum := 0
	maxSum := 0

	for i, v := range arr {
		windowSum += v

		// Since i starts from 0, we have to check if i+1 has reached the windowSize or not.
		// We have to skip the maxSum and windowSum comparison and every other operation until we find the sum of
		// windowSize many elements first.
		if i+1 < windowSize {
			continue
		}

		if maxSum < windowSum {
			maxSum = windowSum
		}

		// remove the first value from the window sum since the window is sliding forward
		windowSum -= arr[i-windowSize+1]
	}
	return maxSum
}
