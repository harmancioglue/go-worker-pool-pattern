package main

import (
	"github.com/harmancioglue/go-worker-pool-pattern/pool"
	"log"
)

func main() {
	totalWorkerCount := 5

	jobList := 15 //job is square of the number
	resultChannel := make(chan int, jobList)

	wP := pool.CreateWorkerPool(resultChannel) //create worker pool
	wP.Handle(totalWorkerCount)                //create workers

	for i := 1; i <= jobList; i++ {
		id := i
		wP.AddItem(id)
	}

	for i := 0; i < jobList; i++ {
		res := <-resultChannel
		log.Printf("Job done. Result: %d", res)
	}
}
