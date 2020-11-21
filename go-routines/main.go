package main

import (
	"fmt"
	"math/rand"
	"time"
)

var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		if b, ok := querryCache(id); ok {
			fmt.Println("From the cache")
			fmt.Println(b)
			continue
		}
		if b, ok := querryDtabase(id); ok {
			fmt.Println("From the database")
			fmt.Println(b)
			continue
		}
		fmt.Println("Book not found")
		time.Sleep(150 * time.Microsecond)
	}
	//println(id)
}
func querryCache(id int) (Book, bool) {
	b, ok := cache[id]
	return b, ok
}
func querryDtabase(id int) (Book, bool) {
	time.Sleep(100 * time.Microsecond)
	for _, b := range books {
		if b.ID == id {
			cache[id] = b
			return b, true
		}
	}
	return Book{}, false
}
