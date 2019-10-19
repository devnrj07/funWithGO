package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Starting the main go routine")
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("-------CPU stats----------")
	fmt.Println("Begin -CPU", runtime.NumCPU())
	fmt.Println("Begin -Go Routines", runtime.NumGoroutine())
	fmt.Println("--------------------------")

	go func() {
		fmt.Println("I'm the one")
		wg.Done()
	}()
	go func() {
		fmt.Println("I'm the second")
		wg.Done()
	}()

	fmt.Println("-------CPU stats----------")
	fmt.Println("Mid -CPU", runtime.NumCPU())
	fmt.Println("Mid -Go Routines", runtime.NumGoroutine())
	fmt.Println("--------------------------")

	wg.Wait()
	fmt.Println("-------CPU stats----------")
	fmt.Println("End -CPU", runtime.NumCPU())
	fmt.Println("End -Go Routines", runtime.NumGoroutine())
	fmt.Println("--------------------------")

	fmt.Println("about to exit")
}
