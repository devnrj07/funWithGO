package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var count int64
	gs := 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&count, 1)
			fmt.Println(atomic.LoadInt64(&count)) // use of atomic to get rid of RACE COndition

		}()
	}

	wg.Wait()
	fmt.Println("End value : ", count)
}
