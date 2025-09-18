package main

import (
	"context"
	"fmt"
	"goRutiiine/miner"
	"goRutiiine/postman"
	"time"
)

func main() {
	mail := make([]string, 0)
	coal := 0

	minerContext, minerCancel := context.WithCancel(context.Background())
	mailContext, mailCancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(3 * time.Second)
		minerCancel()
	}()

	go func() {
		time.Sleep(6 * time.Second)
		mailCancel()
	}()

	coalTransferPoint := miner.MinerPool(minerContext, 2)
	mailTransferPoint := postman.PostmanPool(mailContext, 3)

	isCoalClosed := false
	isMailClosed := false
	for !isCoalClosed || !isMailClosed {
		select {
		case c, ok := <-coalTransferPoint:
			if !ok {
				isCoalClosed = true
				continue
			}
			coal += c
		case m, ok := <-mailTransferPoint:
			if !ok {
				isMailClosed = true
				continue
			}
			mail = append(mail, m)
		}
	}

	fmt.Println("Coal: ", coal, "Mails", len(mail))
}
