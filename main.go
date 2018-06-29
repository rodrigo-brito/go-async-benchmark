package go_async_benchmark

import "time"

func Run(numberTasks int, maxConcurrently int, taskTime time.Duration, runner TaskRunner) {
	tasks := make([]Task, 0, numberTasks)
	for i := 0; i < numberTasks; i ++ {
		tasks = append(tasks, func() {
			time.Sleep(taskTime)
		})
	}
	runner.Run(maxConcurrently, tasks...)
}
