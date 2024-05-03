package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup
var sum = 0

func getGOMAXPROCS() int {
	return runtime.GOMAXPROCS(0)
}

func counter(n int, wg *sync.WaitGroup) {
	time.Now().UnixNano()
	min := 10
	max := 30
	tt := rand.Intn(max-min+1) + min
	for i := 0; i < 10000; i++ {
		sum = sum + 1
		time.Sleep(time.Duration(tt * 1000))

	}
	fmt.Println("From ", n, ":", sum)
	wg.Done()

}

func main() {

	//numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(2)
	fmt.Printf("GOMAXPROCS is %d\n", getGOMAXPROCS())
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go counter(i, &wg)
	}

	wg.Wait()
	fmt.Println("Final Sum:", sum)
}
