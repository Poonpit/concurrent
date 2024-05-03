package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func pingpong(m chan string, wg *sync.WaitGroup) {
	// defer wg.Done()
	msg := <-m
	fmt.Println(msg)
	m <- "pong"
	wg.Done()
}

func main() {
	messages := make(chan string)
	go pingpong(messages, &wg)

	messages <- "ping"
	wg.Add(1)
	msg := <-messages
	fmt.Println(msg)

	wg.Wait()
	fmt.Println("done")
}
