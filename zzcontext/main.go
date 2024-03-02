package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	WithTimeout()
}

func WithTimeout() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 16*time.Second)
	defer cancel()

	wg := &sync.WaitGroup{}

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go work2(ctx, wg, i)
	}

	wg.Wait()
	fmt.Println("done")
}

func work2(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()

	switch id {
	case 0:
		ctx, cancel := context.WithTimeout(ctx, time.Second)
		time.AfterFunc(500*time.Millisecond, func() { cancel() })
		slowTask(ctx, id, fmt.Sprintf("worker %d had a timeout of 1 second", id))
	case 1:
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		slowTask(ctx, id, fmt.Sprintf("worker %d had a timeout of 10 second", id))
	case 2:
		ctx, cancel := context.WithTimeout(ctx, -1*time.Second)
		defer cancel()
		slowTask(ctx, id, fmt.Sprintf("worker %d had a timeout of -1 second", id))
	case 3:
		ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
		defer cancel()
		slowTask(ctx, id, fmt.Sprintf("worker %d had a timeout of 0 second", id))
	}
}

type ID string

func slowTask(ctx context.Context, id int, prefix string) {
	ctx = context.WithValue(ctx, ID("id"), id)

	fmt.Printf("%s: started\n", prefix)

	select {
	case <-time.Tick(15 * time.Second):
		fmt.Printf("%s: finisihed\n", prefix)
	case <-ctx.Done():
		fmt.Printf("%s: too slow... returning %s\n", prefix, ctx.Err())
	}

}

func WithCancel() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	wg := sync.WaitGroup{}

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			work(ctx, i)
			wg.Done()
		}()
	}

	time.AfterFunc(4*time.Second, func() { cancel() })
	wg.Wait()
	log.Println("completed")
}

func work(ctx context.Context, i int) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	slowFn(ctx, i)
}

func slowFn(ctx context.Context, i int) {
	if i == 1 {
		time.Sleep(5 * time.Second)
	}
	log.Printf("Slow function %d started\n", i)
	select {
	case <-time.Tick(3 * time.Second):
		log.Printf("Slow function %d finished\n", i)
	case <-ctx.Done():
		log.Printf("Slow function %d is to slow: %s\n", i, ctx.Err())
	}

}
