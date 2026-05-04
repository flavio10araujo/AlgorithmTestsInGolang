package main

import (
	"fmt"
	"sync"
)

func producers(jobs chan<- int) {
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)
}

func consumer(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Println("Consumer", id, "processing job", job)
		results <- job * 2
	}
}

func main() {
	jobs := make(chan int)
	results := make(chan int)

	var wg sync.WaitGroup

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go consumer(w, jobs, results, &wg)
	}

	go producers(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println("Resultado: ", result)
	}
}
