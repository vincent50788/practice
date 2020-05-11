package main

import (
	"fmt"
	"reflect"
)

type Account struct {
	id      string
	name    string
	balance float64
}

func main() {
	account := Account{
		id:      "X123",
		name:    "Justin Lin",
		balance: 1000,
	}
	t := reflect.TypeOf(account)
	fmt.Println(t.Kind())   // struct
	fmt.Println(t.String()) // main.Account
	/*
	   底下顯示
	   id string
	   name string
	   balance float64
	*/
	for i, n := 0, t.NumField(); i < n; i++ {
		F := t.Field(i)
		fmt.Println(F.Name, F.Type)
		fmt.Printf("%s \n", F.Tag)
	}
}
