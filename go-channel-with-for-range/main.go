package main

import (
	"log"
	"sync"
)

func main() {
	dataChan := make(chan struct{})

	type Counter struct {
		sync.Mutex
		value int
	}

	var counter int
	var waitGroup = new(sync.WaitGroup)

	// create workers
	for i := 1; i <= 100; i++ {
		// consumer
		go func(workerNo int, wg *sync.WaitGroup, c *int, ch chan struct{}) {
			defer wg.Done()
			wg.Add(1)

			for v := range dataChan {
				log.Printf("value #%d: %+v\n", workerNo, v)
				*c++
			}
		}(i, waitGroup, &counter, dataChan)
	}

	// producer
	for i := 0; i < 1000; i++ {
		dataChan <- struct{}{}
	}

	close(dataChan)
	waitGroup.Wait()
	log.Printf("total consumer: %d\n", counter)
	log.Println("terminating...")
}
