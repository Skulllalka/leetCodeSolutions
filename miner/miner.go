package miner

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Miner(ctx context.Context,
	wg *sync.WaitGroup,
	transferPoint chan<- int,
	n int,
	power int) {

	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Miner number %d has ended the worker day", n)
			return
		default:
			fmt.Printf("Miner number %d started to work\n", n)
			time.Sleep(1 * time.Second)
			fmt.Printf("Miner number %d closed work\n", n)
			transferPoint <- power
			fmt.Printf("Miner number %d transferred coal\n", n)
		}

	}
}

func MinerPool(ctx context.Context, minerCount int) <-chan int {
	coalTransferPoint := make(chan int)

	wg := &sync.WaitGroup{}
	for i := 1; i <= minerCount; i++ {
		wg.Add(1)
		go Miner(ctx, wg, coalTransferPoint, i, i*10)
	}

	go func() {
		wg.Wait()
		close(coalTransferPoint)
	}()

	return coalTransferPoint
}
