package main

import (
	"container/ring"
	"fmt"
)

func main() {
	integers := []int{1, 3, 5, 7}
	circular_list := ring.New(len(integers))
	for i := 0; i < circular_list.Len(); i++ {
		circular_list.Value = integers[i]
		circular_list = circular_list.Next()
	}
	circular_list.Do(func(element interface{}) {
		fmt.Print(element, ",")
	})

	fmt.Println()
	for i := 0; i < circular_list.Len(); i++ {
		fmt.Print(circular_list.Value, ",")
		circular_list = circular_list.Prev()
	}
	fmt.Println()
	circular_list = circular_list.Move(2)
	circular_list.Do(func(element interface{}) {
		fmt.Print(element, ",")
	})
	fmt.Println()
}
