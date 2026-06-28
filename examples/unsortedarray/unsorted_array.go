package main

import (
	"fmt"

	"github.com/koktlzz/go-data-structure/arrays"
)

func main() {
	array, err := arrays.NewUnsortedArray[string](5)
	if err != nil {
		fmt.Println(err)
		return
	}
	array.Append("a") // ["a"]
	array.Append("b") // ["a","b"]
	array.Append("c") // ["a","b","c"]
	array.Delete(1)   // ["a","c"]
	array.Find("c")   // 1
	array.Traverse(func(s string) { fmt.Println(s) })
}
