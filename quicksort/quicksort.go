package quicksort

/*
quicksort is a recursive function that sorts -in this case- a slice of integers.
Initial call to the function starts with list itself, starting position of the list
and the last position of the list.

We know that the pivot is sorted when the new position of the pivot is returned from partition function.
We then continue sorting by taking lists left and right of the pivot chosen by using the value/position returned from
partition. This operation continues until "lo" and "hi" values cross.
*/
func quicksort(l []int, lo, hi int) {
	if lo < hi {
		j := partition(l, lo, hi)
		quicksort(l, lo, j-1)
		quicksort(l, j+1, hi)
	}
}

/*
Steps:
	1. Increment i until you find an element greater than pivot
	2. Decrement j until you find an element less than or equal to the pivot
	3. If i < j, swap positions of elements at i and j
	4. When i > j stop and swap pivot with the value at position j. Pivot is now sorted.
*/
func partition(l []int, lo, hi int) int {
	i, j := lo, hi
	mid := (lo + hi) / 2
	pivot := l[mid]

	for i < j {
		// loop until element greater or equal to pivot is found
		for {
			if l[i] > pivot {
				break
			}
			i++
		}

		// loop until element less than pivot is found
		for {
			if j == 0 || l[j] <= pivot {
				break
			}
			j--
		}

		// we know that value at i is bigger than pivot as well as we know that value at j is smaller.
		// swap positions of these values so right side always has greater values while left has smaller ones.
		if i < j {
			l[i], l[j] = l[j], l[i]
		}
	}

	// right position for the pivot is found hence pivot is "sorted"
	l[mid], l[j] = l[j], l[mid]
	return j
}
