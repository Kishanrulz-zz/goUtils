package main

import (
	"fmt"
	"time"
	"os"
	//"os/signal"
	"os/signal"
)

func main() {
	message := make(chan string) // non buffered channel
	count := 3


	gracefulShutdown := make(chan os.Signal, 1)

	signal.Notify(gracefulShutdown)

	go func() {
		<-gracefulShutdown
		time.Sleep(time.Second * 2)
		os.Exit(1)
	}()

	go func() {
		for i := 1; i <= count; i++ {
			fmt.Println("send message")
			message <- fmt.Sprintf("message %d", i)
		}
	}()

	time.Sleep(time.Second * 3)

	for i := 1; i <= count; i++ {
		fmt.Println(<-message)
	}
}
