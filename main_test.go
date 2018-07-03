package go_async_benchmark

import (
	"testing"
	"time"
)

const (
	numberTasks = 1000
	maxAsync = 50
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

func BenchmarkRunnerC(b *testing.B) {
	for n := 0; n < b.N; n++ {
		runner := new(RunnerC)
		Run(numberTasks, maxAsync, duration, runner)
	}
}

func BenchmarkRunnerD(b *testing.B) {
	for n := 0; n < b.N; n++ {
		runner := new(RunnerD)
		Run(numberTasks, maxAsync, duration, runner)
	}
}

func BenchmarkRunnerE(b *testing.B) {
	for n := 0; n < b.N; n++ {
		runner := new(RunnerE)
		Run(numberTasks, maxAsync, duration, runner)
	}
}