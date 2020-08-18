### Insertion Sort Algorithm

### Time Complexity
**Average case:** O(n^2)

**Worst case:** O(n^2)

### Description
Insertion sort is suitable for small data sets. It's also a stable (order of appearance does not change) and in-place algorithm.

The logic is based on comparing the pivot value with the elements on its left side. If there are elements greater than the pivot,
those elements are shifted to right and pivot value is placed in its correct position. 

Steps: 
1. Starting looping from the second element of the list (i=1)
2. Set your pivot to the value at current position of _i_
3. Loop through elements on the left side of your pivot, starting from the one next to the pivot (j = i - 1)
4. Move values one position up to make space for the pivot value
5. Stop when you find an element greater than or equal to the pivot, or reach the beginning of the list
6. Copy the pivot value to the position where you have stopped
7. Increment _i_ and repeat until you reach the end of the list 