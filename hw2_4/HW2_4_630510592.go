package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	str  string
	wait chan bool
}

func say(msg string) <-chan Message { 
	c := make(chan Message)
	waitForIt := make(chan bool)
	go func() { 
		for i := 0; ; i++ {

			c <- Message{fmt.Sprintf("%s: %d", msg, i), waitForIt}
			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
			// <-waitForIt //option
		}
	}()
	return c 
}

func fanIn(input1, input2 <-chan Message) <-chan Message {
	c := make(chan Message)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func main() {
	
	c := fanIn(say("hello"), say("world"))

	for i := 0; i < 4; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)
		// msg1.wait <- true   //option
		// msg2.wait <- true   //option
	}

	fmt.Println("Done")
}