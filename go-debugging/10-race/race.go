package main

import (
	"fmt"
	"sync"
)

// Spot the issue?

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		go func() {
			wg.Add(1)
			fmt.Println(i)
			wg.Done()
		}()
	}

	wg.Wait()
}
