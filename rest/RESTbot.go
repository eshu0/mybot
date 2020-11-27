package main

import (
	"fmt"

	Request "github.com/eshu0/RESTServer/pkg/request"
	"github.com/eshu0/mybot"
)

type Rbot struct {
	mbot *mybot.MyBot
}

func NewRbot() Rbot {
	mbot := mybot.NewMyBot()
	rb := Rbot{mbot: mbot}
	return rb
}

func (bot Rbot) Stop(request Request.ServerRequest) {
	request.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println("Rbot - Stop")
	bot.mbot.Stop()
}

func (bot Rbot) Forwards(request Request.ServerRequest) {
	request.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println("Rbot - Forwards")
	bot.mbot.Forwards()
}

func (bot Rbot) Backwards(request Request.ServerRequest) {
	request.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println("Rbot - Backwards")
	bot.mbot.Forwards()
}

func (bot Rbot) SpinLeft(request Request.ServerRequest) {
	request.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	fmt.Println("Rbot - SpinLeft")
	bot.mbot.SpinLeft()
}

func (bot Rbot) SpinRight(request Request.ServerRequest) {
	request.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println("Rbot - SpinRight")
	bot.mbot.SpinRight()
}

func (bot Rbot) Close() {
	if bot.mbot != nil {
		bot.mbot.Stop()
		bot.mbot.Close()
	}
}
