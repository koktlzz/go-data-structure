package main

import (
	"fmt"

	"github.com/koktlzz/go-data-structure/arrays"
)

func main() {
	array, err := arrays.NewSortedArray[string](5)
	if err != nil {
		fmt.Println(err)
		return
	}
	array.Insert("a")       // ["a"]
	array.Insert("c")       // ["a","c"]
	array.Insert("b")       // ["a","b","c"]
	array.Insert("e")       // ["a","b","c","e"]
	array.Insert("d")       // ["a","b","c","d","e"]
	array.BinarySearch("c") // 2
	array.Delete("c")       // ["a","b","d","e"]

}
