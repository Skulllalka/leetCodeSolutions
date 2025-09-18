package postman

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Postman(ctx context.Context, wg *sync.WaitGroup, transferPoint chan<- string, n int, mail string) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("The day is over, number", n)
			return
		default:
			fmt.Printf("Postman number %d took the mail\n", n)
			time.Sleep(1 * time.Second)
			fmt.Printf("Postman number %d has delivered the mail\n", n)
			transferPoint <- mail
			fmt.Printf("Postman number %d has ended its wor", n)
		}

	}
}

func PostmanPool(ctx context.Context, postmanCount int) <-chan string {
	mailTransferPoint := make(chan string)
	wg := &sync.WaitGroup{}

	for i := 1; i <= postmanCount; i++ {
		wg.Add(1)
		go Postman(ctx, wg, mailTransferPoint, i, postmanToMail(i))
	}

	go func() {
		wg.Wait()
		close(mailTransferPoint)
	}()

	return mailTransferPoint
}

func postmanToMail(postmanNumber int) string {
	ptm := map[int]string{
		1: "hello",
		2: "invite",
		3: "info",
	}

	mail, ok := ptm[postmanNumber]
	if !ok {
		return "lotery"
	}
	return mail
}
