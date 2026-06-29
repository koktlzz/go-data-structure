package main

import (
	"github.com/koktlzz/go-data-structure/arrays"
)

func main() {
	array, _ := arrays.NewDynamicArray[string](2)

	array.Insert("a") // ["a"]
	array.Insert("b") // ["a", "b"]
	array.Insert("c") // ["a", "b", "c"], capacity: 2 -> 4

	array.Find("b")   // 1
	array.Delete("b") // ["a", "c"]
	array.Delete("c") // ["a"], capacity: 4 -> 2
}
