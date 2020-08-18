## Merge Sort Algorithm

### Time Complexity
**Average case:** O(nlogn)

**Worst case:** O(nlogn)

### Description
Merge sort is a divide-and-conquer algorithm. It breaks down the given array/slice into two halves and calls itself recursively
until each sublist has only one element in them. Then, it begins sorting by comparing values of two lists.

The comparing works if and only if the lists are **SORTED**. That's why lists are broken down to single element sublists first,
and comparisons begin from the smallest sublist and works up to larger ones.

When working with arrays/slices, a separate result list is created and values are copied to it.


| Pros  | Cons |
| ---- | ---- |
| Suitable for very large lists.  | Not an in-place sort algorithm, needs extra space (in case of arrays, not linked lists).|
| Suitable for linked lists. | Slow for small sized lists compared to other sort algorithms. |
| Stable; orders of appearance are maintained even if there are doubles | Recursive and like all recursive functions, uses stacks which means needs stack space allocation. |
