##Quicksort Algorithm

###Time Complexity
**Average case:** O(nlogn)

**Worst case:** O(n2) -i.e. when you choose the first element as pivot and it's the smallest value in the list

###Description
Quicksort is a divide-and-conquer algorithm. It can be implemented as in-place without having to need
additional space for the resulting list.

Most of the logic is implemented in the partitioning bit.
During partitioning, a pivot has to be chosen so other values in the list can get compared to it.
When choosing a pivot one can:
- Pick the first element in the list
- Pick the last element in the list
- Pick a random element
- Pick the element in the middle (implemented in quicksort.go in this repo)

After the pivot is chosen, 2 pointers to point the beginning (i), and the end (j) of the list are set. Steps to follow afterwards
are as follows:
1. Increment *i* until you find an element greater than pivot
2. Decrement *j* until you find an element less than or equal to the pivot
3. If values that satisfy conditions at steps 1 and 2 are found, swap them for we want bigger values at the
right side of the list and smaller ones at the left. 
4. When i > j stop and swap pivot with the value at position *j*. Pivot is now **sorted**.
5. Repeat steps above for lists left and right of the pivot's position.