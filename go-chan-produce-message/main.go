package main

import (
	"log"
	"time"
)

type ChanGroup struct {
	ResultChan        chan int
	ResultUpdatedChan chan bool
	ExitChan          chan bool
}

func main() {
	g := ChanGroup{
		ResultChan:        make(chan int),
		ResultUpdatedChan: make(chan bool),
		ExitChan:          make(chan bool),
	}

	go func() {
		// exit
		time.Sleep(20 * time.Second)
		g.ExitChan <- true
	}()

	go func() {
		// produce something
		var count int
		for {
			count++
			time.Sleep(2 * time.Second)
			log.Printf("Send Result: %d", count)
			g.ResultChan <- count

			log.Println("Wait Confirm Process Result...")
			<-g.ResultUpdatedChan
			log.Println("Confirm Process Result")
		}
	}()

	go func() {
		// consume something
		for {
			log.Printf("Receive Result: %d", g.GetResult())
			g.ProcessResultDone()
		}
	}()

	<-g.ExitChan
	log.Println("Exit Program")
}

func (g ChanGroup) ProcessResultDone() {
	g.ResultUpdatedChan <- true
}

func (g ChanGroup) GetResult() int {
	return <-g.ResultChan
}
