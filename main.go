package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func processData(v int) int {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	return v * 2
}
func main() {
	in := make(chan int)
	out := make(chan int)

	go func() {
		for i := range 10 {
			in <- i
		}
		close(in)
	}()
	start := time.Now()
	processThread(in, out, 5)

	for v := range out {
		fmt.Println("v=", v)
	}
	fmt.Println("main duration:", time.Since(start))
}

func processThread(in <-chan int, out chan<- int, threadsNumber int) {
	wg := &sync.WaitGroup{}
	for range threadsNumber {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range in {
				out <- processData(v)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()
}
