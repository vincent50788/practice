package main

import (
	"log"
	_ "log"
)

//Car struct
type Car struct {
	model  Model
	engine Engine
	tire   Tire
}

//Info 規格
type Info struct {
	brand string
	model string
}

//Model 車種 接口---------------------------------------
type Model interface {
	info()
}

//RAV4 is a kind of model
type RAV4 struct {
	brand string
	model string
}

//info can get model RAV4's information
func (m RAV4) info() {
	m.brand = "Toyota"
	m.model = "RAV4"
	log.Printf("%s %s 啟動\n", m.brand, m.model)
}

//Engine 引擎種類 接口-----------------------------------
type Engine interface {
	info()
}

//V8 is an kind of Engine
type V8 struct {
	brand string
	model string
}

//info can get engine V8's information
//info can get model RAV4's information
func (m V8) info() {
	m.brand = "Toyota"
	m.model = "V8"
	log.Printf("引擎:%s\n", m.model)
}

//Tire 輪台種類 接口-------------------------------------
type Tire interface {
	info()
}

//Continental is a kind of Tire
type Continental struct {
	brand string
	model string
}

//info can get Tire Continental's information
func (m Continental) info() {
	m.brand = "Continental"
	log.Printf("輪胎:%s\n", m.brand)
}

//Start the car
func (car Car) Start() {
	car.model.info()
	car.engine.info()
	car.tire.info()
}

func main() {
	toyota := &Car{
		model:  new(RAV4),
		engine: new(V8),
		tire:   new(Continental),
	}
	toyota.Start()
}

// func LogInit() {
// 	log.SetFlags(log.Ldate | log.Lshortfile)
// }
