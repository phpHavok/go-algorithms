package sort

// MergeSort performs the merge sort algorithm on a slice of integers and
// returns a new slice containing those integers in ascending order.
func MergeSort(numbers []int) []int {
	// The singleton and empty lists are both already sorted.
	if len(numbers) < 2 {
		return numbers
	}
	// Split the list into halves and sort the halves.
	mid := len(numbers) / 2
	left := MergeSort(numbers[:mid])
	right := MergeSort(numbers[mid:])
	// Then, combine the sorted halves with a helper function.
	return mergeSlices(left, right)
}

// mergeSlices combines two sorted (ascending) slices of integers into a single
// slice of integers that is sorted between the two.
func mergeSlices(left, right []int) []int {
	// Make a new slice with the capacity of the combined lengths of the input
	// slices.
	size := len(left) + len(right)
	merged := make([]int, 0, size)
	var rest []int
	// Fill the slice with sorted values from each of the input slices, taking
	// the smallest element from either as we move on. If one slice runs out of
	// elements, the rest of the remaining slice (already in sorted order by
	// definition) should be appended onto the end of the larger slice.
	for k := 0; k < size; k++ {
		if len(left) < 1 {
			rest = right
			break
		}
		if len(right) < 1 {
			rest = left
			break
		}
		if left[0] <= right[0] {
			merged = append(merged, left[0])
			left = left[1:]
		} else {
			merged = append(merged, right[0])
			right = right[1:]
		}
	}
	merged = append(merged, rest...)
	return merged
}
