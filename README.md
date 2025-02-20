# Go Slice Append Bug
This example demonstrates a common misunderstanding of how the `append` function works with slices in Go when the capacity is exceeded.

## The Bug
The `append` function in Go will efficiently add elements to a slice if there is enough capacity. If the capacity is exceeded, it will allocate a new slice with a larger capacity and copy all existing elements into it, before adding the new elements. This can lead to unexpected performance implications if not handled carefully.
The provided code demonstrates this by creating a slice with a pre-allocated capacity of 10, then appending more than 10 elements. After the 10th element, the slice's underlying array is reallocated resulting in two distinct memory blocks.

## The Solution
The solution involves either allocating a slice with a sufficiently large initial capacity to avoid exceeding it during appends or using techniques like dynamically increasing capacity to efficiently handle large number of appends.

## How to Reproduce
1. Save the code in `bug.go`.
2. Run the code using `go run bug.go`.
3. Observe the output. The memory addresses of the slice will change after the capacity is exceeded.

## How to Solve
Refer to the `bugSolution.go` for the improved implementation.