package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

var mu sync.Mutex
var sum = 0
var csum chan int

func counter(n int, wg *sync.WaitGroup, csum chan int) {

	for i := 0; i < 10000; i++ {
		sum = <-csum
		sum = sum + 1
		go func() { csum <- sum }()

	}

	fmt.Println("Form ", n, ":", sum)
	wg.Done()

}

func main() {
	csum := make(chan int)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go counter(i, &wg, csum)
	}
	csum <- 0
	wg.Wait()
	sum = <-csum
	fmt.Println("Final sum:", sum)
	close(csum)
}
