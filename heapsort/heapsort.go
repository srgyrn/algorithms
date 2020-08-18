package heapsort

func heapsort(l []int) {
	j := len(l) - 1
	heapify(l, 0, j)

	for i := len(l) - 1; i > 0; i-- {
		// Move the largest element to the end of the list and the last child to the root.
		l[0], l[j] = l[i], l[0]
		// Sort the list
		heapify(l, 0, i-1)
		j--
	}
}

// heapify builds a max heap
func heapify(heap []int, start, size int) {
	for i := size / 2; i >= start; i-- {
		posLeftChild := (2 * i) + 1
		posRightChild := (2 * i) + 2

		// Check if the current element is a parent node, continue if not
		if posLeftChild > size {
			continue
		}

		posGreaterChild := posLeftChild

		// Check if the node has a right child. If it does, compare the children and take the bigger value
		if posRightChild <= size && heap[posLeftChild] < heap[posRightChild] {
			posGreaterChild = posRightChild
		}

		// If parent is smaller than its biggest child's value, swap them
		if heap[i] < heap[posGreaterChild] {
			heap[i], heap[posGreaterChild] = heap[posGreaterChild], heap[i]

			// Recursively check and heapify the affected subtree
			heapify(heap, posGreaterChild, size)
		}
	}
}
