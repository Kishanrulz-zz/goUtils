package main

import "fmt"

//func sendData(sendch chan<- int) {
//	sendch <- 10
//}
//
//func main() {
//	sendch := make(chan<- int)
//	go sendData(sendch)
//	fmt.Println(<-sendch)
//}

//usage of send only channel
func sendData(sendch chan<- int) {
	sendch <- 10
}

func main() {
	chnl := make(chan int)
	go sendData(chnl)
	fmt.Println(<-chnl)
}