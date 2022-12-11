package main

import (
	"fmt"
	"sync"
	"time"
)

var JOBS int = 10    // How many jobs each version will schedule
var BATCHES int = 10 // How large each batch will be for concurrency

func calculate(id string) {
	/**
	* Desc: Function simulates a heavy calculation task
	*
	* Args:
	*	id (int): An id for the task
	*		Ex. Sequential (1/10)
	*			Concurrent (4/5)
	 */
	fmt.Printf("Starting %s..\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s done..\n", id)
}

func sequential(total int) {
	/**
	* Desc: Function simulates n sequential calls to calculate
	*
	* Args:
	*	total (int): Total amount of jobs
	 */
	for i := 0; i < total; i++ {
		calculate(fmt.Sprintf("Sequential %d/%d", i, total))
	}
}

func worker(id int, total int, batch_size int, wg *sync.WaitGroup) {
	/**
	* Desc: Worker in a sync.WaitGroup during concurrent execution
	*
	* Args:
	*	id (int): An id for the worker
	*	total (int): Total amount of jobs
	*	batch_size (int): How many tasks the worker is responsible for
	*	wg (*sync.WaitGroup): The waitgroup for the worker
	 */
	defer wg.Done()
	for i := 0; i < batch_size; i++ {
		calculate(fmt.Sprintf("Concurrent %d/%d", id*batch_size+i, total))
	}
}

func concurrent(total int, batches int) {
	/**
	* Desc: Function that simulates n calls divided to m workers
	*		and executed concurrently
	*
	* Args:
	*	total (int): How many calls in total should be simulated
	*	batches (int): How many workers should divide the work
	 */
	var wg sync.WaitGroup
	for i := 0; i < batches; i++ {
		wg.Add(1)
		go worker(i, total, total/batches, &wg)
	}

	wg.Wait()
}

func main() {
	fmt.Println("Golang Runtimes:")

	start := time.Now()
	sequential(JOBS)
	duration := time.Since(start)
	fmt.Printf("Sequential: %f s\n", duration.Seconds())

	start = time.Now()
	concurrent(JOBS, BATCHES)
	duration = time.Since(start)
	fmt.Printf("Concurrent: %f s\n", duration.Seconds())
}
