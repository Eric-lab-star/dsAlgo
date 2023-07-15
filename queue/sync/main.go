package main

// importing fmt package
import (
	"fmt"
	"math/rand"
	"time"
)

// constants
const (
	messagePassStart = iota
	messageTicketStart
	messagePassEnd
	messageTicketEnd
)

// Queue class
type Queue struct {
	waitPass    int
	waitTicket  int
	playPass    bool
	playTicket  bool
	queuePass   chan int
	queueTicket chan int
	message     chan int
}

// New method initialises queue
func (queue *Queue) New() {

	queue.message = make(chan int)
	queue.queuePass = make(chan int)
	queue.queueTicket = make(chan int)

	go func() {
		var message int
		for {
			select {
			case message = <-queue.message:
				switch message {
				case messagePassStart:
					queue.waitPass++
				case messagePassEnd:
					queue.playPass = false
				}
				if queue.waitPass > 0 && !queue.playPass {
					queue.playPass = true
					queue.waitPass--
					queue.queuePass <- 1
				}
			}
		}
	}()
}

func (queue *Queue) StartPass() {
	queue.message <- messagePassStart
	<-queue.queuePass
}

func (queue *Queue) EndPass() {
	queue.message <- messagePassEnd
}

func passenger(queue *Queue) {
	queue.StartPass()
	queue.EndPass()

}

// main method
func main() {
	var queue *Queue = &Queue{}
	queue.New()
	for i := 0; i < 10; i++ {
		go passenger(queue)
	}
	time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
	fmt.Println("end")
}
