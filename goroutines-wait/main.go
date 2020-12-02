package main

import (
	"context"
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
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		go func(ctx context.Context) {
			for {
				select {
				case <-ctx.Done():
					log.Println("inside goroutine stop")
					return
				case <-time.Tick(1 * time.Second):
					log.Println("inside goroutine leak")
				}
			}
		}(ctx)
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
	time.Sleep(1 * time.Second)
	log.Printf("num=%d", runtime.NumGoroutine())
	log.Println("the end")
}
