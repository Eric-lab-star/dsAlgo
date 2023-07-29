package main

import (
	"fmt"
)

func PrintSpiral(n int) []int {
	size := n * n
	top := 0
	bottom := n - 1
	left := 0
	right := n - 1
	slice := make([]int, size)
	i := 0
	for left < right {
		for c := left; c <= right; c++ {
			slice[top*n+c] = i
			i++
		}
		top++

		for r := top; r <= bottom; r++ {
			slice[n*r+right] = i
			i++
		}
		right--

		for c := right; c >= left; c-- {
			slice[bottom*n+c] = i
			i++
		}
		bottom--

		for r := bottom; r >= top; r-- {
			slice[r*n+left] = i
			i++
		}
		left++
	}
	slice[top*n+left] = i

	return slice

}

func main() {
	n := 5
	length := 2
	for i, sketch := range PrintSpiral(n) {
		fmt.Printf("%*d ", length, sketch)
		if i%n == n-1 {
			fmt.Println("")
		}
	}
}
