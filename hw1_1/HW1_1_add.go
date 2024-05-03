package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

var mu sync.Mutex
var sum = 0
// var csum chan int

func counter(n int, wg *sync.WaitGroup) {

	for i := 0; i < 10000; i++ {
		mu.Lock()
		sum = sum + 1
		mu.Unlock()
	}
	fmt.Println("Form ", n, ":", sum)
	wg.Done()
	// mu.Unlock()
}

func main() {

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go counter(i, &wg)
	}
	wg.Wait()
	fmt.Println("Final sum:", sum)
}
