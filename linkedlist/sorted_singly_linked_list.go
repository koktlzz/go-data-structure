package linkedlist

import "cmp"

type SortedSinglyLinkedList[T cmp.Ordered] struct {
	head *node[T]
}

func (ssll *SortedSinglyLinkedList[T]) InsertInSortedList(newData T) {
	cur := ssll.head
	var pre *node[T]
	for cur != nil {
		if cur.data >= newData {
			if pre == nil {
				ssll.head = newNode(newData, cur) // Add the element at the beginning of the list
			} else {
				pre.append(newNode(newData, cur)) // General Case
			}
			return
		}
		pre = cur
		cur = cur.next
	}
	if pre == nil {
		ssll.head = newNode(newData, nil) // The list is empty
	} else {
		pre.append(newNode(newData, nil)) // Add the element at the end of the list
	}
}
