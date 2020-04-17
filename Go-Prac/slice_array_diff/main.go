package main

import (
	"log"
	"unsafe"
)

func main(){
	s := make([]byte, 200)
	ptr := unsafe.Pointer(&s[0])
	log.Print("ptr: ", &ptr)
}