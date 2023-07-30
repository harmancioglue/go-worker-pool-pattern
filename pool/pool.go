package pool

import (
	"fmt"
	"time"
)

type IWorkerPool interface {
	Handle()
	AddItem(func())
}

type workerPool struct {
	queue  chan int
	result chan int
}

func CreateWorkerPool(resultChannel chan int) *workerPool {
	fmt.Printf("Worker pool created")
	return &workerPool{
		queue:  make(chan int),
		result: resultChannel,
	}
}

func (wP *workerPool) Handle(workerCount int) {
	for i := 1; i <= workerCount; i++ {
		fmt.Printf("Worker %d created\n", i)
		go func(id int) {
			for job := range wP.queue {
				time.Sleep(3 * time.Second)
				fmt.Println(fmt.Sprintf("Worker %d process job", id))
				wP.result <- job * job
			}
		}(i)
	}
}

func (wP *workerPool) AddItem(job int) {
	wP.queue <- job
}
