package linkedlist

import "fmt"

type SinglyLinkedList[T comparable] struct {
	head *node[T]
}

type node[T comparable] struct {
	data T
	next *node[T]
}

func newNode[T comparable](data T, nextNode *node[T]) *node[T] {
	return &node[T]{
		data: data,
		next: nextNode,
	}
}

func (n *node[T]) append(nextNode *node[T]) {
	n.next = nextNode
}

func NewSinglyLinkedList[T comparable]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{}
}

func (sll *SinglyLinkedList[T]) InsertInFront(data T) {
	sll.head = &node[T]{
		data: data,
		next: sll.head,
	}
}

func (sll *SinglyLinkedList[T]) InsertToBack(data T) {
	cur := sll.head
	tail := newNode(data, nil)
	if cur == nil {
		sll.head = tail
		return
	}
	for cur.next != nil {
		cur = cur.next
	}
	cur.append(tail)
}

func (sll *SinglyLinkedList[T]) Delete(target T) error {
	cur := sll.head
	var pre *node[T]
	for cur != nil {
		if cur.data == target {
			if pre == nil {
				sll.head = cur.next
				return nil
			} else {
				pre.append(cur.next)
				return nil
			}
		}
		pre = cur
		cur = cur.next
	}
	return fmt.Errorf("No element with value %v was found in the list.", target)
}

func (sll *SinglyLinkedList[T]) DeleteFromFront() error {
	if sll.head == nil {
		return fmt.Errorf("Delete on an empty list.")
	}
	sll.head = sll.head.next
	return nil
}
