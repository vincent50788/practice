package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	ch := make(chan string, 1)
	defer func() {
		close(ch)
	}()

	go func() {
		for {
			time.Sleep(1 * time.Second)
			fmt.Println(time.Now().Unix())
			i++

			select {
			case ch <- "100":
				for v := range ch {
					fmt.Printf("%v\n" ,v + string(i))
				}
			}
		}
	}()

	time.Sleep(time.Second * 1)
	ch <- "stop"
}