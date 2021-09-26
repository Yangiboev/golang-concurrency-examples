package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("worker ", id, "started job", job)
		time.Sleep(time.Second)
		fmt.Println("worker ", id, "finished job", job)
		results <- job * 2
	}
}
func main() {
	const numberOfJobs = 3
	const numberOfWorkers = 3
	jobs := make(chan int, numberOfJobs)
	results := make(chan int, numberOfJobs)

	for w := 1; w <= numberOfWorkers; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numberOfJobs; j++ {
		jobs <- j
	}
	close(jobs)
	for i := 0; i < numberOfJobs; i++ {
		<-results
	}
	close(results)

}
