package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	broadcastChan := make(chan struct{}, 100)
	wg := new(sync.WaitGroup)
	wg.Add(5)
	go func(ch <-chan struct{}) {
		defer wg.Done()
		select {
		case v, ok := <-ch:
			log.Printf("receive chan #1: %v %v", v, ok)
		}
	}(broadcastChan)
	go func(ch <-chan struct{}) {
		defer wg.Done()
		select {
		case v, ok := <-ch:
			log.Printf("receive chan #2: %v %v", v, ok)
		}
	}(broadcastChan)
	go func(ch <-chan struct{}) {
		defer wg.Done()
		select {
		case v, ok := <-ch:
			log.Printf("receive chan #3: %v %v", v, ok)
		}
	}(broadcastChan)
	go func(ch <-chan struct{}) {
		defer wg.Done()
		select {
		case v, ok := <-ch:
			log.Printf("receive chan #4: %v %v", v, ok)
		}
	}(broadcastChan)
	go func(ch <-chan struct{}) {
		defer wg.Done()
		select {
		case v, ok := <-ch:
			log.Printf("receive chan #5: %v %v", v, ok)
		}
	}(broadcastChan)

	time.Sleep(1 * time.Second)
	close(broadcastChan)
	wg.Wait()
	log.Printf("terminating...")
}
