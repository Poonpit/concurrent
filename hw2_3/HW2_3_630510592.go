package main

import (
	"fmt"
	"sync"
)

func say(wg *sync.WaitGroup, i int) {
	fmt.Println("World: ", i)
	wg.Done() 
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 4; i++ {
		fmt.Println("Hello: ", i)
		wg.Add(1) 
		go say(&wg, i)
		wg.Wait()
	}

	fmt.Println("Done")
}
