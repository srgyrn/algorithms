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

// SmallestSubArrayWithGivenSum finds the given sum in the smallest contiguous subarray
func SmallestSubArrayWithGivenSum(arr []int, sum int) int {
	var s int
	var windowStart int
	var minCount int

	for i, v := range arr {
		s += v

		// if s exceeds sum, extract the value at windowStart and windowStart forward until s is less than sum again
		for s > sum {
			s -= arr[windowStart]
			windowStart++
		}

		if s == sum {
			count := i - windowStart + 1
			if minCount == 0 || count < minCount {
				minCount = count
			}
			s = v
			windowStart = i
		}
	}

	return minCount
}


// LongestSubstringWithKDistinct finds the longest contiguous subarray with k distinct characters
func LongestSubstringWithKDistinct(s string, k int) (int, string) {
	var maxCount int
	var result string
	var windowStart int
	charFrequency := make(map[rune]int, k+1)

	for i, v := range s {
		if _, ok := charFrequency[v]; !ok {
			charFrequency[v] = 1
		} else {
			charFrequency[v]++
		}

		if len(charFrequency) > k {
			count := i - windowStart
			if maxCount < count {
				maxCount = count
				result = s[windowStart:i]
			}

			for len(charFrequency) >= k {
				index := rune(s[windowStart])
				charFrequency[index]--
				if c := charFrequency[index]; c == 0 {
					delete(charFrequency, index)
				}
				windowStart++
			}
		}
	}

	return maxCount, result
}