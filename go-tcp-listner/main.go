package main

import (
	"context"
	"log"
	"net"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", ":1000")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	go func() {
		for {
			_, err := ln.Accept()
			if err != nil {
				// check error
				oe, ok := err.(*net.OpError)
				if ok && oe.Temporary() {
					continue
				}
				if ok && oe.Timeout() {
					continue
				}

				log.Println(err)
				return
			}
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	select {
	case <-ctx.Done():
		ln.Close()
	}
	time.Sleep(1 * time.Second)
	log.Println("Exit")
}
