package main


//func main() {
//	ch := make(chan int)
//	ch <- 5
//}

//func main() {
//	c := make(chan int)
//	c <- 42    // write to a channel
//	val := <-c // read from a channel
//	println(val)
//}

func main() {
	c := make(chan int)

	// Make the writing operation be performed in
	// another goroutine.
	go func() {
		c <- 42
	}()
	val := <-c
	println(val)
}