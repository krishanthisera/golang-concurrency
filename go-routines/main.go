package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	wg := &sync.WaitGroup{}
	mx := &sync.RWMutex{}
	for i := 0; i < 10; i++ {
		wg.Add(2)
		id := rnd.Intn(10) + 1
		go func(id int, wg *sync.WaitGroup, mx *sync.RWMutex) {
			if b, ok := querryCache(id, mx); ok {
				fmt.Println("From the cache")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg, mx)
		go func(id int, wg *sync.WaitGroup, mx *sync.RWMutex) {
			if b, ok := querryDtabase(id, mx); ok {
				fmt.Println("From the database")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg, mx)
		//fmt.Println("Book not found")
		//time.Sleep(150 * time.Microsecond)
	}
	//println(id)
	wg.Wait()
}
func querryCache(id int, mx *sync.RWMutex) (Book, bool) {
	mx.RLock()
	b, ok := cache[id]
	mx.RUnlock()
	return b, ok
}
func querryDtabase(id int, mx *sync.RWMutex) (Book, bool) {
	//time.Sleep(100 * time.Microsecond)
	for _, b := range books {
		if b.ID == id {
			mx.Lock()
			cache[id] = b
			mx.Unlock()
			return b, true
		}
	}
	return Book{}, false
}
