package main

import (
	"fmt"
	"time"

	Request "github.com/eshu0/RESTServer/pkg/request"
	"github.com/eshu0/mybot"
)

type Rbot struct {
	mbot *mybot.MyBot
}

func NewRbot(folder string) Rbot {
	mbot := mybot.NewMyBot(folder)
	rb := Rbot{mbot: mbot}
	return rb
}

func (bot Rbot) Stop(request Request.ServerRequest) {
	// stupid CORS - Means that for dev this is a pain
	// should be removed for PROD or for internet access
	request.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println("Rbot - Stop")
	bot.mbot.Stop()
}

func (bot Rbot) Forwards(request Request.ServerRequest) {
	// stupid CORS - Means that for dev this is a pain
	// should be removed for PROD or for internet access
	request.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println("Rbot - Forwards")
	bot.mbot.Forwards()
}

func (bot Rbot) Backwards(request Request.ServerRequest) {
	// stupid CORS - Means that for dev this is a pain
	// should be removed for PROD or for internet access
	request.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println("Rbot - Backwards")
	bot.mbot.Backwards()
}

func (bot Rbot) TurnLeft(request Request.ServerRequest) {
	request.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	fmt.Println("Rbot - TurnLeft")
	bot.mbot.SpinLeft()
	time.Sleep(200 * time.Millisecond)
	bot.mbot.Stop()

}

func (bot Rbot) TurnRight(request Request.ServerRequest) {
	// stupid CORS - Means that for dev this is a pain
	// should be removed for PROD or for internet access
	request.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println("Rbot - TurnRight")
	bot.mbot.SpinRight()
	time.Sleep(200 * time.Millisecond)
	bot.mbot.Stop()
}

func (bot Rbot) SpinLeft(request Request.ServerRequest) {
	// stupid CORS - Means that for dev this is a pain
	// should be removed for PROD or for internet access
	request.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	fmt.Println("Rbot - SpinLeft")
	bot.mbot.SpinLeft()
}

func (bot Rbot) SpinRight(request Request.ServerRequest) {
	// stupid CORS - Means that for dev this is a pain
	// should be removed for PROD or for internet access
	request.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println("Rbot - SpinRight")
	bot.mbot.SpinRight()
}

func (bot Rbot) Capture(request Request.ServerRequest) {
	// stupid CORS - Means that for dev this is a pain
	// should be removed for PROD or for internet access
	request.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println("Rbot - Capture")
	s, err := bot.mbot.Capture()

	if err == nil {
		fmt.Println("Rbot - Capture Success")
		fmt.Println(s)
	} else {
		fmt.Println("Rbot - Error:")
		fmt.Println(err.Error())
	}
}

func (bot Rbot) Close() {
	if bot.mbot != nil {
		bot.mbot.Stop()
		bot.mbot.Close()
	}
}
