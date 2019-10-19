package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	count := 0
	gs := 100
	var wg sync.WaitGroup
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {

			temp := count
			runtime.Gosched()
			temp++
			count = temp
			fmt.Println("counter value ->", count)
			wg.Done()

		}()
	}

	wg.Wait()
	fmt.Println("End value : ", count)
}
