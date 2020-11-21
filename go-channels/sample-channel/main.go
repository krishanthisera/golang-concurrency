package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 5)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func(w *sync.WaitGroup, c <-chan int) {
		for msg := range ch {
			fmt.Println(msg)
		}
		wg.Done()
	}(wg, ch)
	go func(w *sync.WaitGroup, c chan<- int) {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}(wg, ch)
	wg.Wait()
}
