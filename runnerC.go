package go_async_benchmark

import "sync"

type RunnerC struct {}

func (r * RunnerC) Run(limit int, tasks ...Task) {
	wg := new(sync.WaitGroup)
	for _, task := range tasks {
		wg.Add(1)
		go func() {
			task()
			wg.Done()
		}()
	}
	wg.Wait()
}
