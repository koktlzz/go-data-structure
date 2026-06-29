package linkedlist

import "fmt"

type DoublyLinkedList[T comparable] struct {
	head *dnode[T]
	tail *dnode[T]
}

type dnode[T comparable] struct {
	data T
	next *dnode[T]
	pre  *dnode[T]
}

func newdNode[T comparable](data T) *dnode[T] {
	return &dnode[T]{
		data: data,
	}
}

func (n *dnode[T]) append(nextNode *dnode[T]) {
	n.next = nextNode
	if nextNode != nil {
		nextNode.pre = n
	}
}

func (n *dnode[T]) prepend(preNode *dnode[T]) {
	n.pre = preNode
	if preNode != nil {
		preNode.next = n
	}
}

func (dll *DoublyLinkedList[T]) InsertInFront(data T) {
	oldHead := dll.head
	if oldHead == nil {
		dll.head = newdNode(data)
		dll.tail = dll.head
	} else {
		dll.head = newdNode(data)
		dll.head.append(oldHead)
	}
}

func (dll *DoublyLinkedList[T]) InsertToBack(data T) {
	oldTail := dll.tail
	if oldTail == nil {
		dll.head = newdNode(data)
		dll.tail = dll.head
	} else {
		dll.tail = newdNode(data)
		dll.tail.prepend(oldTail)
	}
}

func (dll *DoublyLinkedList[T]) Delete(target T) error {
	cur := dll.head
	for cur != nil {
		if cur.data == target {
			if cur.pre == nil {
				dll.head = cur.next // Delete node from front
				if dll.head == nil {
					dll.tail = nil // In case the list's head was the only element in the list
				} else {
					dll.head.pre = nil
				}
			} else if cur.next == nil {
				dll.tail = cur.pre  // Delete node from back
				dll.tail.next = nil // We know tail can't be None, because cur.pre != nil
			} else {
				cur.pre.append(cur.next) // General Case
				cur.next.prepend(cur.pre)
			}
			return nil
		}
		cur = cur.next
	}
	return fmt.Errorf("No element with value %v was found in the list.", target)
}
