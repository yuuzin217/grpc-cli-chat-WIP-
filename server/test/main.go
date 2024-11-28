package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int, 3)
	go func(ch chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			time.Sleep(time.Second * 1)
			ch <- 3
		}
	}(ch, &wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			time.Sleep(time.Second * 2)
			fmt.Println("len:", len(ch))
			// fmt.Println("received:", <-ch)
		}
	}(ch, &wg)
	wg.Wait()
}
