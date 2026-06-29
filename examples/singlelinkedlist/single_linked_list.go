package main

import "github.com/koktlzz/go-data-structure/linkedlist"

func main() {
	list := linkedlist.NewSingleLinkedList[string]()

	list.InsertToBack("b")  // ["b"]
	list.InsertInFront("a") // ["a","b"]
	list.InsertToBack("c")  // ["a","b","c"]
	list.Delete("b")        // ["a","c"]
	list.DeleteFromFront()  // ["c"]
}
