package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2)
	go one(&wg)
	
	go two(&wg)
	wg.Wait()
}

func one(wg *sync.WaitGroup) {
	fmt.Println("I'm the one")
	wg.Done()
}

func two(wg *sync.WaitGroup) {

	fmt.Println("I'm the two")
	wg.Done()
}
