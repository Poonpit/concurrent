package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func player(wg *sync.WaitGroup, name string, action chan string, win chan string) {

	for {
		msg := <-action

		if msg == "Hit" {
			wg.Done()
			return
		}

		n := rand.Intn(20)
		if n == 0 {
			action <- "Hit" 
			fmt.Println(name,">> Hit")
			win <- name
			// wg.Done()
			return
		}

		if msg == "Ping" {
			fmt.Print(name, ": Ping \n")
			action <- "Pong"
		} else {
			fmt.Print(name, ": Pong \n")
			action <- "Ping"
		}

	}

}

func main() {
	var wg sync.WaitGroup

	action := make(chan string)
	win := make(chan string)


	wg.Add(1)
	go player(&wg, "Alice", action, win)
	go player(&wg, "Bob", action, win)
	action <- "Ping"

	wg.Wait()
	whowin := <- win
	fmt.Println("Game over")
	fmt.Println(whowin, "win")
	close(action)
	close(win)
}