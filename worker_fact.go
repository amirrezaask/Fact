package main

import "math/big"

type factJob struct {
	upperbound *big.Int
	lowerbound *big.Int
}

func worker(c <-chan *factJob, res chan<- *big.Int) {
	for j := range c {
		res <- fact(j.upperbound, j.lowerbound)
	}
}

func createFactWorkers(num int, jobs <-chan *factJob, results chan<- *big.Int) {
	for i := 0; i < num; i++ {
		go worker(jobs, results)
	}
}

//WorkerFact calculates a factorial using a worker pool
func WorkerFact(n int) {
	jobs := make(chan *factJob, 10)
	results := make(chan *big.Int, 10)

}
