package arrays

import (
	"fmt"
)

type UnsortedArray[T comparable] struct {
	arr     []T
	size    int
	maxSize int
}

func NewUnsortedArray[T comparable](maxSize int) (*UnsortedArray[T], error) {
	if maxSize < 0 {
		return nil, fmt.Errorf("maxSize cannot be negative: %d.", maxSize)
	}
	return &UnsortedArray[T]{
		arr:     make([]T, maxSize),
		size:    0,
		maxSize: maxSize,
	}, nil
}

func (ua *UnsortedArray[T]) Append(newEntry T) error {
	if ua.size >= ua.maxSize {
		return fmt.Errorf("The array is already full.")
	}
	ua.arr[ua.size] = newEntry
	ua.size++
	return nil
}

func (ua *UnsortedArray[T]) Delete(index int) error {
	if ua.size == 0 {
		return fmt.Errorf("Delete from an empty array.")
	} else if index < 0 || index >= ua.size {
		return fmt.Errorf("Index(%d) out of range.", index)
	} else {
		ua.arr[index] = ua.arr[ua.size-1]
		ua.size--
		return nil
	}
}

func (ua *UnsortedArray[T]) Find(target T) int {
	for i := range ua.size {
		if ua.arr[i] == target {
			return i
		}
	}
	return -1
}

func (ua *UnsortedArray[T]) Traverse(callBack func(T)) {
	for i := range ua.size {
		callBack(ua.arr[i])
	}
}
