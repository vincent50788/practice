package main

import (
	"fmt"
)

type NEET struct {
	phone Phone
	game  Game
}

type BadNEET struct {
	phone *piexl4
	game  *gameboy
}

type Game interface {
	play()
}

type ps4 struct {
}

func (p ps4) play() {
	fmt.Println("play ps4")
}

type gameboy struct {
}

func (p gameboy) play() {
	fmt.Println("play gameboy")
}

type Phone interface {
	call(param int) string
	takephoto()
}

type piexl4 struct {
}

func (p piexl4) call(param int) string {
	fmt.Println("i am piexl4, i can call you!", param)
	return "damon"
}

func (p piexl4) takephoto() {
	fmt.Println("i can take a photo for you")
}

type iphone struct {
}

func (i iphone) call(param int) string {
	fmt.Println("this is iphone", param)
	return string(param)
}

func (i iphone) takephoto() {
	fmt.Println("shooting photos with my iphone")
}

func main() {
	/*
		japanNeetBad := &BadNEET{
			phone: new(piexl4),
			game:  new(gameboy),
		}
		japanNeetBad.game.play()
		japanNeetBad.phone.call(1)
		japanNeetBad.phone.takephoto()
	*/

	//==============

	japanNeet := &NEET{
		phone: new(piexl4),
		game:  new(gameboy),
	}
	japanNeet.game.play()
	japanNeet.phone.call(1)
	japanNeet.phone.takephoto()

	japanNeet.game = new(ps4)
	japanNeet.phone = new(iphone)
	japanNeet.phone.call(1)
	japanNeet.phone.takephoto()
	japanNeet.game.play()
}
