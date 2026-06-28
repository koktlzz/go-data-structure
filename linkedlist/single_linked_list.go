package linkedlist

import "fmt"

type SingleLinkedList[T comparable] struct {
	head *node[T]
}

type node[T comparable] struct {
	data T
	next *node[T]
}

func newNode[T comparable](data T) *node[T] {
	return &node[T]{
		data: data,
	}
}

func (n *node[T]) append(nextNode *node[T]) {
	n.next = nextNode
}

func NewSingleLinkedList[T comparable]() *SingleLinkedList[T] {
	return &SingleLinkedList[T]{}
}

func (sll *SingleLinkedList[T]) InsertInFront(data T) {
	sll.head = &node[T]{
		data: data,
		next: sll.head,
	}
}

func (sll *SingleLinkedList[T]) InsertToBack(data T) {
	cur := sll.head
	tailNode := newNode(data)
	if cur == nil {
		sll.head = tailNode
		return
	}
	for cur.next != nil {
		cur = cur.next
	}
	cur.append(tailNode)
}

func (sll *SingleLinkedList[T]) Delete(target T) error {
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

func (sll *SingleLinkedList[T]) DeleteFromFront() error {
	if sll.head == nil {
		return fmt.Errorf("Delete on an empty list.")
	}
	sll.head = sll.head.next
	return nil
}
