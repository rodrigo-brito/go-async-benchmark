package go_async_benchmark

import (
	"sync"

)

type RunnerB struct {}

func (r * RunnerB) TaskConsummer(tasks chan Task, wg *sync.WaitGroup) {
	for {
		task, more := <-tasks
		if !more {
			return
		}

		task()
		wg.Done()
	}
}


func (r * RunnerB) Run(limit int, tasks ...Task) {
	wg := new(sync.WaitGroup)
	tasksChannel := make(chan Task)

	for i := 0; i < limit; i++ {
		go r.TaskConsummer(tasksChannel, wg)
	}

	for i := 0; i < len(tasks); i++ {
		wg.Add(1)
		tasksChannel <- tasks[i]
	}
	close(tasksChannel)
	wg.Wait()
}
