package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()

	dataChan := make(chan struct{})

	type Counter struct {
		sync.Mutex
		value int
	}
	var counter = new(Counter)
	var waitGroup = new(sync.WaitGroup)

	// create workers
	for i := 1; i <= 100; i++ {
		// consumer
		go func(workerNo int, wg *sync.WaitGroup, c *Counter, ch chan struct{}) {
			defer wg.Done()
			wg.Add(1)

			for v := range dataChan {
				log.Printf("value #%d: %+v\n", workerNo, v)
				c.Lock()
				c.value++
				c.Unlock()
			}
		}(i, waitGroup, counter, dataChan)
	}

	// producer
	for i := 0; i < 1000; i++ {
		dataChan <- struct{}{}
	}

	close(dataChan)
	waitGroup.Wait()

	log.Printf("elapsed time: %s\n", time.Since(startTime))
	log.Printf("total consumer: %d\n", counter)
	log.Println("terminating...")
}
