package main

import (
	"context"
	"fmt"
	"time"
)

type AppCtx struct {
	State int
}

type AppCtxKey int

const (
	appValueCtxKey AppCtxKey = iota
)

func main() {
	bgCtx := context.Background()

	appStateCtx, cancel := context.WithCancel(bgCtx)
	defer cancel()

	appCtx := AppCtx{
		State: 1,
	}
	appValueCtx := context.WithValue(appStateCtx, appValueCtxKey, &appCtx)

	go func(ctx context.Context) {
		appCtx := ctx.Value(appValueCtxKey).(*AppCtx)
		fmt.Printf("%+v\n", appCtx)
		appCtx.State++
		fmt.Printf("%+v\n", appCtx)

		select {
		case <-ctx.Done():
			fmt.Println("Done!")
			return
		}
	}(appValueCtx)

	go func() {
		time.Sleep(1 * time.Second)
		cancel()
	}()

	select {
	case <-appStateCtx.Done():
		time.Sleep(1 * time.Second)
		fmt.Printf("%+v\n", appCtx)
		fmt.Println("Exit!")
		return
	}
}
