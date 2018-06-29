package go_async_benchmark

type RunnerA struct {}

func (r * RunnerA) Run(limit int, tasks ...Task) {
	for _, task := range tasks {
		task()
	}
}
