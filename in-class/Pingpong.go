package main

import (
	"fmt"
	"time"
	"sync"
)

func pingpong(messages chan string) {
	for {
		msg := <-messages
		fmt.Println("Bob:", msg)
		time.Sleep(time.Second) 
		messages <- "pong"
	}
}

func main() {
	var wg sync.WaitGroup
	messages := make(chan string)

	go pingpong(messages)

	for i := 0; i < 3; i++ {
		time.Sleep(time.Second) 
		messages <- "ping"
		result := <-messages
		fmt.Println("Alice:", result)
	}
	// time.Sleep(2 * time.Second) 
	wg.Wait()
	close(messages)
	fmt.Println("Alice:", "hit!!!")
}
