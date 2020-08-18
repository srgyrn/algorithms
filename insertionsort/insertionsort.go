package insertionsort

func insertionSort(l []int) {
	for i := 1; i < len(l); i++ {
		pivot := l[i]
		j := i - 1

		// shift values on the left of the pivot until you reach a value bigger than the pivot
		// or you reach the beginning of the list
		for j >= 0 && pivot < l[j] {
			l[j + 1] = l[j]
			j--
		}

		// place the pivot in correct position
		l[j + 1] = pivot
	}
}
