package main

import (
	"github.com/harmancioglue/go-worker-pool-pattern/pool"
	"log"
	"time"
)

func main() {
	totalWorkerCount := 5
	wP := pool.CreateWorkerPool() //create worker pool
	wP.Handle(totalWorkerCount)   //create workers

	totalTask := 5 //task is square of the number
	resultC := make(chan int, totalTask)

	for i := 1; i <= totalTask; i++ {
		id := i
		wP.AddItem(func() {
			log.Printf("Starting task %d", id)
			time.Sleep(5 * time.Second)
			resultC <- id * 2
		})
	}

	for i := 0; i < totalTask; i++ {
		res := <-resultC
		log.Printf("Task has been finished with result %d:", res)
	}

}
