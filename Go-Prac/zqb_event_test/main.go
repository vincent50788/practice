package main

import (
	"gitlab.paradise-soft.com.tw/glob/eventbus"
	"log"
	"sync"
	"time"
)

type CallbackFunc = eventbus.CallbackFunc
type EntryID = eventbus.EntryID

var e *event

var wg sync.WaitGroup

type Topic int

const (
	A Topic = iota
	B
	C
	D
	E
	F
	G

)

type event struct {
	eventbus.IEvent
}

func New() {
	e = &event{
		IEvent: eventbus.New(eventbus.WithInternalEngine()),
	}
}

func (e *event) Subscribe(topic Topic, callback CallbackFunc) (EntryID, error) {
	return e.IEvent.Subscribe(string(topic), callback)
}

func (e *event) Publish(topic Topic, message interface{}) error {
	return e.IEvent.Publish(string(topic), message)
}

func Unmarshal(data []byte, key interface{}) error {
	return eventbus.Unmarshal(data, key)
}

func main() {
	New()

	wg.Add(1)
	go func(){
		for i:=0; i<100; i++{
			e.Publish(A, i)
			time.Sleep(time.Millisecond * 500)
		}
		wg.Done()
	}()

	time.Sleep(time.Second * 3)

	id, err := e.Subscribe(A, callback)
	if err != nil{
		log.Print(err)
	}

	log.Print("id: ", id)

	wg.Wait()
}

func callback (data []byte) {
	var s int
	err :=  Unmarshal(data, &s)
	if err != nil {
		log.Print(err)
	}

	log.Print("s: ", s)
}