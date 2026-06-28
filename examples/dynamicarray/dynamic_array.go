package main

import (
	"fmt"

	"github.com/koktlzz/go-data-structure/arrays"
)

func main() {
	array, err := arrays.NewDynamicArray[string](2)
	if err != nil {
		fmt.Println(err)
		return
	}

	array.Insert("a") // ["a"]
	array.Insert("b") // ["a", "b"]
	array.Insert("c") // ["a", "b", "c"], capacity: 2 -> 4

	array.Find("b")   // 1
	array.Delete("b") // ["a", "c"]
	array.Delete("c") // ["a"], capacity: 4 -> 2
}
