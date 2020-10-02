## Sliding Window Technique

### Time Complexity
**Average case:** O(n)

### Description
Sliding window technique is best for solving problems when a max or min calculated value is expected from a contiguous
subset. 

The aim is to avoid calculating overlapping operations. Hence, a subset of elements are taken (window) and as it moves 
forward, operations (adding, subtracting etc.) are applied to the next item while extracting the value of the item
that is no longer in the window.