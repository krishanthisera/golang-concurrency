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
	//Creating a waight group
	//Waitgroup will keep the process alive until the execution of the subrotine finished
	wg := &sync.WaitGroup{}
	//Mutexes allow a paticular operation to do changes on a resource
	mx := &sync.RWMutex{}

	//Chennels for cache and database
	dbCh := make(chan Book)
	cacheCh := make(chan Book)
	for i := 0; i < 10; i++ {

		wg.Add(2) //Define Number of wait groups
		id := rnd.Intn(10) + 1

		/*************************************************************************
		*Routine 01
		*************************************************************************/
		go func(id int, wg *sync.WaitGroup, mx *sync.RWMutex, ch chan<- Book) {
			if b, ok := querryCache(id, mx); ok {
				ch <- b
			}
			wg.Done()
		}(id, wg, mx, cacheCh)

		/*************************************************************************
		*Routine 02
		*************************************************************************/
		go func(id int, wg *sync.WaitGroup, mx *sync.RWMutex, ch chan<- Book) {
			if b, ok := querryDtabase(id, mx); ok {
				mx.Lock()
				cache[id] = b
				mx.Unlock()
				ch <- b
			}
			wg.Done()
		}(id, wg, mx, dbCh)

		/*************************************************************************
		*Routine 03
		*************************************************************************/
		go func(cacheCh, dbCh chan Book) {
			select {
			case b := <-cacheCh:
				fmt.Println("Value from cache!!")
				fmt.Println(b)
				<-dbCh
			case b := <-dbCh:
				fmt.Println("Value from database!!")
				fmt.Println(b)
			}
		}(cacheCh, dbCh)

	}

	//println(id)
	wg.Wait()
	close(dbCh)
	close(cacheCh)

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
