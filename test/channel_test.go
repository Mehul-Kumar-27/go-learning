package main

import (
	channelsExample "conversion/channels"
	"fmt"
	"testing"
	"time"
)

func TestPayment(t *testing.T) {
	payment := channelsExample.NewPayment()
	go payment.Start()
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		go func(i int) {
			payment.UserChan <- fmt.Sprintf("User %d", i)

		}(i)

		time.Sleep(time.Millisecond * 1000)
	}
}
