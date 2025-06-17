package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID int
}

type Result struct {
	JobID int
	Value int
}

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		// Simulate work
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("Worker %d processed job %d\n", id, job.ID)
		results <- Result{JobID: job.ID, Value: job.ID * 2}
	}
}

func main() {
	numWorkers := 3
	numJobs := 5

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	var wg sync.WaitGroup

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j}
	}
	close(jobs)

	// Wait for workers to finish, then close results
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	for res := range results {
		fmt.Printf("Result: job %d -> value %d\n", res.JobID, res.Value)
	}
}
