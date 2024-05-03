package main

import "fmt"


func Bob(ch chan string){
	msg := <- ch
	if (msg == "ping"){
		ch <- "pong"
	}else {
		ch <- "ping"
	}
}

func Alice(ch chan string){
	msg := <- ch
	if (msg == "ping"){
		ch <- "pong"
	}else {
		ch <- "ping"
	}
}

func main() {
	ch := make(chan string)
	ch <- "ping"

	// go func ()  {
	// 	messages <- "ping"
	// }()

	// msg := <- messages
	// fmt.Println(msg)
	go Bob(ch)
	go Alice(ch)
	
	fmt.Println("Pong")
}