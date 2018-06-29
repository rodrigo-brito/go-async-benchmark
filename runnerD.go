package go_async_benchmark

import "sync"

type RunnerD struct {}

func (r * RunnerD) Run(limit int, tasks ...Task) {
	wg := new(sync.WaitGroup)
	control := make(chan bool, limit)

	for _, task := range tasks {
		wg.Add(1)
		control <- true

		go func() {
			task()
			<-control
			wg.Done()
		}()
	}
	wg.Wait()
}