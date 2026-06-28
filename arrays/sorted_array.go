package arrays

import (
	"cmp"
	"fmt"
)

type SortedArray[T cmp.Ordered] struct {
	arr     []T
	size    int
	maxSize int
}

func NewSortedArray[T cmp.Ordered](maxSize int) (*SortedArray[T], error) {
	if maxSize < 0 {
		return nil, fmt.Errorf("maxSize cannot be negative: %d.", maxSize)
	}
	return &SortedArray[T]{
		arr:     make([]T, maxSize),
		size:    0,
		maxSize: maxSize,
	}, nil
}

func (sa *SortedArray[T]) Insert(value T) error {
	if sa.size >= sa.maxSize {
		return fmt.Errorf("The array is already full.")
	}
	for i := sa.size; i > 0; i-- {
		if sa.arr[i-1] <= value {
			sa.arr[i] = value
			sa.size++
			return nil
		} else {
			sa.arr[i] = sa.arr[i-1]
		}
	}
	sa.arr[0] = value
	sa.size++
	return nil
}

func (sa *SortedArray[T]) LinearSearch(target T) int {
	for i := range sa.size {
		if sa.arr[i] == target {
			return i
		} else if sa.arr[i] > target {
			return -1
		}
	}
	return -1
}

func (sa *SortedArray[T]) BinarySearch(target T) int {
	left, right := 0, sa.size-1
	for left <= right {
		mid := (left + right) / 2
		midVal := sa.arr[mid]
		if midVal == target {
			return mid
		} else if midVal > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func (sa *SortedArray[T]) Delete(target T) error {
	index := sa.BinarySearch(target)
	if index == -1 {
		return fmt.Errorf("Unable to delete element %v: the entry is not in the array", target)
	} else {
		for i := index; i < sa.size-1; i++ {
			sa.arr[i] = sa.arr[i+1]
		}
		sa.size--
		return nil
	}
}
