package pool

import "fmt"

type IWorkerPool interface {
	Handle()
	AddItem(func())
}

type workerPool struct {
	queue chan func()
}

func CreateWorkerPool() *workerPool {
	fmt.Printf("Worker pool created")
	return &workerPool{
		queue: make(chan func()),
	}
}

func (wP *workerPool) Handle(workerCount int) {
	for i := 1; i <= workerCount; i++ {
		fmt.Printf("Worker %d created\n", i)
		go func(id int) {
			for task := range wP.queue {
				fmt.Println(fmt.Sprintf("Worker %d process task", id))
				task()
			}
		}(i)
	}
}

func (wP *workerPool) AddItem(task func()) {
	wP.queue <- task
}
