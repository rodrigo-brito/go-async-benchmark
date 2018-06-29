package go_async_benchmark

import (
	"testing"
	"time"
)

const (
	numberTasks = 1000
	maxAsync = 10
	duration = 10 * time.Millisecond
)


func BenchmarkRunnerA(b *testing.B) {
	for n := 0; n < b.N; n++ {
		runner := new(RunnerA)
		Run(numberTasks, maxAsync, duration, runner)
	}
}


func BenchmarkRunnerB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		runner := new(RunnerB)
		Run(numberTasks, maxAsync, duration, runner)
	}
}
