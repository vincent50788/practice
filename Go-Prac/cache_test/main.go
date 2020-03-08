package main

import (
	"errors"
	"fmt"
	"gitlab.paradise-soft.com.tw/glob/cache"
	"log"
	"time"

	"github.com/ahmetb/go-linq"
)

const (
	cars = "car"
)

var Cache *cache.Cache

type Person struct {
	Name string
}

type Car struct {
	year int
	owner, model string
}

type PaymentApp struct {
	ID          int32       `sql:"id"`           // 識別碼
	Name        string      `sql:"name"`         // 名稱
	Image       string      `sql:"image"`        // 圖片路徑
	Status      int32       `sql:"status"`       // 啟用狀態0=禁用1=啟用
}


func main() {
	//ExampleQuery_ToMap()
	//res := map[int32]*PaymentApp{}
	Cache := cache.NewCache(-1, 0)
	paymentAppCache := make(map[int32]*PaymentApp)
	paymentAppCache[1] = &PaymentApp{
		ID:     1,
		Name:   "aaa",
		Image:  "aaa",
		Status: 1,
	}
	paymentAppCache[2] = &PaymentApp{
		ID:     2,
		Name:   "bbb",
		Image:  "bbb",
		Status: 1,
	}
	paymentAppCache[3] = &PaymentApp{
		ID:     3,
		Name:   "ccc",
		Image:  "ccc",
		Status: 1,
	}
	Cache.Set(cars, paymentAppCache, -1)
	cacheResults, _:= Cache.Get(cars)
	log.Println(cacheResults)

	fruits := []string{"apple", "passionfruit", "banana", "mango",
		"orange", "blueberry", "grape", "strawberry"}
	var query []string
	linq.From(fruits).
		SelectT(func(i interface{})string{
		return 	i.(string)
	}).Where(
			func(i interface{}) bool {
				return len(i.(string)) > 6
			},
		).
		Where(
			func(i interface{}) bool {return i.(string) == "grape"},
			).
		ToSlice(&query)

	fmt.Println(query)
}

func GetPerson() (interface{}, error) {
	return &Person{"Adam"}, nil
}

func GetPersonErr() (interface{}, error) {
	return &Person{"Adam"}, errors.New("unknown error")
}

func GetPerson2() (interface{}, error) {
	return &Person{"Adam" + time.Now().String()}, nil
}

func ExampleQuery_ToMap() {
	type Product struct {
		Name string
		Code int
	}

	products := []Product{
		{Name: "orange", Code: 4},
		{Name: "apple", Code: 9},
		{Name: "lemon", Code: 12},
		{Name: "apple", Code: 9},
	}

	map1 := map[int]string{}
	linq.From(products).
		SelectT(
			func(item Product) linq.KeyValue { return linq.KeyValue{Key: item.Code, Value: item.Name} },
		).
		ToMap(&map1)

	log.Println(map1)
	//fmt.Println(map1[4])
	//fmt.Println(map1[9])
	//fmt.Println(map1[12])
	// Output:
	// orange
	// apple
	// lemon
}