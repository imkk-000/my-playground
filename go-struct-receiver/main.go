package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type bag struct {
	number int32
}

func (b *bag) Counter() {
	atomic.AddInt32(&b.number, 1)
}

func (b bag) Number() int32 {
	return b.number
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan struct{}, 1)
	bch := make(chan struct{}, 1)
	b := bag{}
	go func() {
		defer wg.Done()
		wg.Add(1)
		for range ch {
			b.Counter()
			bch <- struct{}{}
		}
		close(bch)
	}()
	fmt.Printf("main: %p\n", &b.number)
	for i := 0; i < 10; i++ {
		ch <- struct{}{}
		fmt.Println(i, b.Number())
		<-bch
	}

	close(ch)
	wg.Wait()
}
