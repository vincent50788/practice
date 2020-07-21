package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func fib() func() int {
	x, y := 0, 1
	return func() int {
		tmp := x
		x, y = y, x+y
		return tmp
	}
}

func main() {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		log.Print(err)
	}

	m := f.(map[string]interface{})
	for k, v := range m {
		switch vt := v.(type) {
		case string:
			fmt.Println(k, "is string", vt)
		case float64:
			fmt.Println(k, "is float64", vt)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vt {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}
