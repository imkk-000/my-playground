package main

import (
	"context"
	"log"
	"time"
)

func main() {
	bgCtx := context.Background()
	ctx1, cancel1 := context.WithTimeout(bgCtx, 1*time.Millisecond)
	defer cancel1()
	ctx2, cancel2 := context.WithTimeout(bgCtx, 2*time.Second)
	defer cancel2()
	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			log.Println("go routine #1 is done")
			return
		}
	}(ctx1)

	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			log.Println("go routine #2 is done")
			return
		}
	}(ctx2)

	time.Sleep(5 * time.Second)
}
