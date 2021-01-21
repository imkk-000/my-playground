package main

import (
	"context"
	"time"
)

type AppContext struct {
	context.Context
}

func main() {
	mainCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx := &AppContext{
		Context: mainCtx,
	}

	myFunc := func(ctx context.Context) {
		select {
		case <-ctx.Done():
			return
		}
	}
	myFunc(ctx)

	go func(cancelFunc context.CancelFunc) {
		select {
		case <-time.After(1 * time.Second):
		}
		cancel()
	}(cancel)
}
