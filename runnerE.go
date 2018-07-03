package go_async_benchmark

import "sync"

type RunnerE struct {}

func runTask(wg *sync.WaitGroup, task Task) {
	task()
	wg.Done()
}

func (r * RunnerE) Run(limit int, tasks ...Task) {
	wg := new(sync.WaitGroup)

	for _, task := range tasks {
		wg.Add(1)
		go runTask(wg, task)
	}

	wg.Wait()
}