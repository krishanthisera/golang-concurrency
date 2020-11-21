package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 5)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func(w *sync.WaitGroup, c chan int) {
		//time.Sleep(5 * time.Second)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		wg.Done()
	}(wg, ch)
	go func(w *sync.WaitGroup, c chan int) {
		c <- 91
		c <- 22
		wg.Done()
	}(wg, ch)
	wg.Wait()
}
