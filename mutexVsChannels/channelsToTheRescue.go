package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(4)

	ch := make(chan bool)

	for i:=1; i<10; i++ {
		for j:=1; j<10;j++ {
			ch <- true
			go func() {
				fmt.Printf("%d + %d = %d\n", i, j, i+j)
				<-ch
			}()
		}
	}
	fmt.Scanln()
}

//
