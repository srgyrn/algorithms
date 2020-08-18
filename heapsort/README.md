## Heap Sort Algorithm

### Time Complexity
**Average case:** O(nlogn)

**Worst case:** O(nlogn)

### Description
Heap sort is an in-place, comparison based sorting algorithm. It works with a complete binary tree and has 2 steps:
1. Build a binary heap (max heap is implemented in this repo)
2. Delete all elements from the heap

##### Steps of creating a max heap:
1. Start from the end and loop until you find a parent node
2. Compare parent's children and take the greatest value
3. Compare the greatest child of the node with the node itself
4. Swap if the node is smaller than its greatest child
5. Take the subtree starting from the position of the greatest child and make a recursive call to make sure
the new parent is at the correct position
6. Repeat process with the next parent

##### Steps of deleting nodes of a max heap:
1. Always delete the root node because it has the largest value
2. Swap values of the root and the last item in the array
3. Heapify the array from position 0 to (len(array) - 1)
4. Repeat until every element is deleted/moved to the correct position

| Node Relations | |
| --- | --- |
| Node | i  | 
| Left child | (2 * i) + 1 |
| Right child | (2 * i) + 2 |
| Parent | floor((i - 1) / 2) |