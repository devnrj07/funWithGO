package main

import (
	"fmt"
	"sync"
)

func main() {

	count := 0
	gs := 100
	var wg sync.WaitGroup
	wg.Add(gs)

	var m sync.Mutex
	for i := 0; i < gs; i++ {
		go func() {
			m.Lock()
			temp := count

			temp++
			count = temp
			fmt.Println("counter value ->", count) //reading an incrementing value causes race condition here. So, we use mutex to lock it
			m.Unlock()                             // now no two go routines will be reading this value at the same time
			wg.Done()

		}()
	}

	wg.Wait()
	fmt.Println("End value : ", count)
}
