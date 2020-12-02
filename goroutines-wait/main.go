package main

import (
	"log"
	"runtime"
	"sync"
	"time"
)

func main() {
	exitChan := make(chan bool, 1)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("inside goroutine")
		go func() {
			for range time.Tick(1 * time.Second) {
				log.Println("inside goroutine leak")
			}
		}()
		time.Sleep(5 * time.Second)
		exitChan <- true
		log.Println("go off")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-time.Tick(1 * time.Second):
				log.Printf("num=%d", runtime.NumGoroutine())
			case <-exitChan:
				log.Println("num off")
				return
			}
		}
	}()

	log.Printf("num=%d", runtime.NumGoroutine())
	log.Println("waiting...")
	wg.Wait()
	time.Sleep(5 * time.Second)
	log.Printf("num=%d", runtime.NumGoroutine())
	log.Println("the end")
}
