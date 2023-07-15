package main

import "fmt"

type Stack []int

func (stack *Stack) Push(prop ...int) {
	*stack = append(*stack, prop...)
}

func (stack *Stack) Pop() int {
	length := len(*stack)
	value := (*stack)[length-1]
	*stack = (*stack)[:length-1]
	return value
}

func main() {
	stack := &Stack{}
	stack.Push(1)
	fmt.Println(stack)
	v := stack.Pop()
	fmt.Println(v)
	fmt.Println(stack)
}
