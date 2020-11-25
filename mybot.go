package mybot

import (
	"fmt"

	"github.com/stianeikeland/go-rpio"
)

type MyBot struct {
	Pin7  rpio.Pin
	Pin11 rpio.Pin
	Pin13 rpio.Pin
	Pin15 rpio.Pin
}

func NewMyBot() *MyBot {

	mbot := &MyBot{}
	fmt.Println("opening gpio")
	err := rpio.Open()
	if err != nil {
		panic(fmt.Sprint("unable to open gpio", err.Error()))
	}

	// 4 is 7
	mbot.Pin7 = rpio.Pin(4)
	mbot.Pin7.Output()

	// 11 is 17
	mbot.Pin11 = rpio.Pin(17)
	mbot.Pin11.Output()

	//pin13 is 27
	mbot.Pin13 = rpio.Pin(27)
	mbot.Pin13.Output()

	// pin 15 is 22
	mbot.Pin15 = rpio.Pin(22)
	mbot.Pin15.Output()

	return mbot

}

func (bot *MyBot) Forwards() {
	fmt.Println("Forwards!")
	bot.Pin15.High()
	bot.Pin11.High()
	bot.Pin13.Low()
	bot.Pin7.Low()
}

func (bot *MyBot) Backwards() {
	fmt.Println("Backwards!")
	bot.Pin15.Low()
	bot.Pin11.Low()
	bot.Pin13.High()
	bot.Pin7.High()
}

func (bot *MyBot) Stop() {
	fmt.Println("Ending....")
	bot.Pin15.Low()
	bot.Pin11.Low()
	bot.Pin13.Low()
	bot.Pin7.Low()
}

func (bot *MyBot) TurnRight() {
	fmt.Println("Turning Right!")

	bot.Pin15.High()
	bot.Pin11.Low()
	bot.Pin13.Low()
	bot.Pin7.High()
}

func (bot *MyBot) TurnLeft() {
	fmt.Println("Turning Left!")
	bot.Pin15.Low()
	bot.Pin11.High()
	bot.Pin13.High()
	bot.Pin7.Low()
}
