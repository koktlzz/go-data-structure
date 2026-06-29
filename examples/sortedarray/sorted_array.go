package main

import (
	"github.com/koktlzz/go-data-structure/arrays"
)

func main() {
	array, _ := arrays.NewSortedArray[string](5)

	array.Insert("a")       // ["a"]
	array.Insert("c")       // ["a","c"]
	array.Insert("b")       // ["a","b","c"]
	array.Insert("e")       // ["a","b","c","e"]
	array.Insert("d")       // ["a","b","c","d","e"]
	array.BinarySearch("c") // 2
	array.Delete("c")       // ["a","b","d","e"]

}
