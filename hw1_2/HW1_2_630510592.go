package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup
var sumChan = make(chan int)
var resultChan = make(chan int)
var finalSumChan = make(chan int)

func getGOMAXPROCS() int {
	return runtime.GOMAXPROCS(0)
}

func counter(n int, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Now().UnixNano()
	min := 10
	max := 30
	tt := rand.Intn(max-min+1) + min
	localSum := 0

	for i := 0; i < 10000; i++ {
		localSum = localSum + 1
		time.Sleep(time.Duration(tt * 1000))
	}

	sumChan <- localSum
}

func main() {
	runtime.GOMAXPROCS(2)
	fmt.Printf("GOMAXPROCS is %d\n", getGOMAXPROCS())

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go counter(i, &wg)
	}

	go func() {
		wg.Wait()
		close(sumChan)
	}()

	go func() {
		totalSum := 0
		for partialSum := range sumChan {
			totalSum += partialSum
			resultChan <- totalSum
		}
		close(resultChan)
	}()

	go func() {
		for i := 0; i < 5; i++ {
			finalSum := <-resultChan
			finalSumChan <- finalSum
		}
		close(finalSumChan)
	}()

	for i := 0; i < 5; i++ {
		finalSum := <-finalSumChan
		fmt.Println("From", i, ":", finalSum)
	}
}
