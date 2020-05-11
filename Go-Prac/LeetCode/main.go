package main

import "fmt"

func fib () func() int{
	x, y := 0, 1
	return func() int {
		tmp := x
		x, y = y, x + y
		return tmp
	}
}


func main(){
	f := fib()
	for i:=0; i<=10; i++{
		fmt.Printf("%v: %v\n", i, f())
	}
}