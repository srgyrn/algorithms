package mergesort

/*
mergseSort divides the given slice in half and sorts the left half, then the right half.

*/
func mergeSort(l []int) []int {
	// if the slice has only one element, return it
	if len(l) < 2 {
		return l
	}

	// divide the slice from half and call itself
	mid := len(l) / 2
	return merge(mergeSort(l[:mid]), mergeSort(l[mid:]))
}

/* merge function takes two slices of the same kind and compares values one by one to create a single sorted list.
For this comparison to work, both slices given should be SORTED slices.
The result is created in a separate slice, so it needs extra space.
*/
func merge(a, b []int) []int {
	i, j, k := 0, 0, 0
	sortedList := make([]int, len(a)+len(b), len(a)+len(b))

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			sortedList[k] = a[i]
			i++
		} else {
			sortedList[k] = b[j]
			j++
		}
		k++
	}

	// copy remaining elements if any left in list a
	for ; i < len(a); i++ {
		sortedList[k] = a[i]
		k++
	}

	// copy remaining elements if any left in list b
	for ; j < len(b); j++ {
		sortedList[k] = b[j]
		k++
	}

	return sortedList
}
