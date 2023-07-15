package main

import "fmt"

type Order struct {
	priority     int
	quantity     int
	product      string
	customerName string
}

type Queue []*Order

func (queue *Queue) Add(order *Order) {
	if len(*queue) == 0 {
		*queue = append(*queue, order)
	} else {
		appended := false
		for i, addedOrder := range *queue {
			if order.priority > addedOrder.priority {
				*queue = append((*queue)[:i], append(Queue{order}, (*queue)[i:]...)...)
				appended = true
				break
			}
		}
		if !appended {
			*queue = append(*queue, order)
		}

	}
}

func (queue *Queue) Iterate() {
	for _, order := range *queue {
		fmt.Println(*order)
	}
}

func main() {
	queue := &Queue{}
	order1 := &Order{2, 20, "Computer", "Greg White"}
	order2 := &Order{1, 10, "Monitoer", "John Smith"}
	order3 := &Order{5, 10, "Monitoer", "John Smith"}
	queue.Add(order1)
	queue.Add(order2)
	queue.Add(order3)
	queue.Iterate()
}
