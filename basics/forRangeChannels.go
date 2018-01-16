package main

import (
	"fmt"
)

//func producer(chnl chan int) {
//	for i := 0; i < 10; i++ {
//		chnl <- i
//	}
//	close(chnl)
//}
//func main() {
//	ch := make(chan int)
//	go producer(ch)
//	for {
//		v, ok := <-ch
//		if ok == false {
//			break
//		}
//		fmt.Println("Received ", v, ok)
//	}
//}

//for range version of the channel

//func producer(chnl chan int) {
//	for i := 0; i < 10; i++ {
//		chnl <- i
//	}
//	close(chnl)
//}
//func main() {
//	ch := make(chan int)
//	go producer(ch)
//	for v := range ch {
//		fmt.Println("Received ",v)
//	}
//}

//Closing a channel

//Unfortunately this code does not work now. Range will work until the channel is closed explicitly.
//All we have to do is to close the channel with a close function
//func main() {
//	message := make(chan string)
//	count := 3
//
//	go func() {
//		for i := 1; i <= count; i++ {
//			message <- fmt.Sprintf("message %d", i)
//		}
//	}()
//
//	for msg := range message {
//		fmt.Println(msg)
//	}
//}

func main() {
	message := make(chan string)
	count := 3

	go func() {
		for i := 1; i <= count; i++ {
			message <- fmt.Sprintf("message %d", i)
		}
		close(message)
	}()

	for msg := range message {
		fmt.Println(msg)
	}
}

