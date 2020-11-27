package main

import (
	"fmt"

	Request "github.com/eshu0/RESTServer/pkg/request"
	"github.com/eshu0/mybot"
)

type Rbot struct {
	mbot mybot.MyBot
}

func NewRbot() Rbot {
	mbot := mybot.NewMyBot()
	rb := Rbot{mbot: mbot}
	return rb
}

func (bot Rbot) Stop(request Request.ServerRequest) {
	fmt.Println("Rbot - Stop")
	bot.mbot.Stop()
}

func (bot Rbot) Forwards(request Request.ServerRequest) {
	fmt.Println("Rbot - Forwards")
	bot.mbot.Forwards()
}

func (bot Rbot) Backwards(request Request.ServerRequest) {
	fmt.Println("Rbot - Backwards")
	bot.mbot.Forwards()
}

func (bot Rbot) SpinLeft(request Request.ServerRequest) {
	fmt.Println("Rbot - SpinLeft")
	bot.mbot.SpinLeft()
}

func (bot Rbot) SpinRight(request Request.ServerRequest) {
	fmt.Println("Rbot - SpinRight")
	bot.mbot.SpinRight()
}

func (bot Rbot) Close(request Request.ServerRequest) {
	fmt.Println("Rbot - Close")
	bot.mbot.Close()
}
