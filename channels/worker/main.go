package main

import (
	"fmt"
	"time"
)

func worker(id int, job <-chan int, result chan<- int) {
	for j := range job {
		fmt.Println("worker %d started %d job", id, j)
		time.Sleep(time.Second)
		fmt.Println("worker %d finished %d job", id, j)
		result <- 2 * j
	}
}

func main() {
	const n = 5
	job := make(chan int, n)
	result := make(chan int, n)

	for w := 1; w <= 3; w++ {
		go worker(w, job, result)
	}

	for j := 1; j <= n; j++ {
		job <- j
	}
	close(job)

	for a := 1; a <= n; a++ {
		fmt.Println(<-result)
	}
}
