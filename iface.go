package go_async_benchmark

type Task func()

type TaskRunner interface {
	Run(limit int, tasks ...Task)
}