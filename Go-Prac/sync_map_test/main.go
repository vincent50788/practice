package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var table sync.Map

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			if _, ok := table.Load("KEY"); !ok {
				table.Store("KEY", n)
			}
		}(i)
	}
	wg.Wait()
	val, ok := table.Load("KEY")
	fmt.Printf("val: %v, ok: %v", val, ok)
}
