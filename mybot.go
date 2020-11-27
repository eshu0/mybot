package mybot

import (
	"fmt"

	"os/exec"
	"path/filepath"
	"time"

	"github.com/stianeikeland/go-rpio"
)

const (
	STILL      = "raspistill"
	HFLIP      = "-hf"
	VFLIP      = "-vf"
	OUTFLAG    = "-o"
	FILE_TYPE  = ".jpg"
	TIME_STAMP = "2006-01-02_15:04:05"
)

type Camera struct {
	horizontalFlip bool
	verticalFlip   bool
	savePath       string
}

type MyBot struct {
	Pin7   rpio.Pin
	Pin11  rpio.Pin
	Pin13  rpio.Pin
	Pin15  rpio.Pin
	CStill *Camera
}

func New(path string) *Camera {
	if path == "" {
		return nil
	}
	return &Camera{false, false, path}
}

func NewMyBot(folder string) *MyBot {

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

	c := camera.New(folder)
	mbot.CStill = c
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
	fmt.Println("Stop!")
	bot.Pin15.Low()
	bot.Pin11.Low()
	bot.Pin13.Low()
	bot.Pin7.Low()
}

func (bot *MyBot) SpinRight() {
	fmt.Println("Spinning Right!")
	bot.Pin15.High()
	bot.Pin11.Low()
	bot.Pin13.Low()
	bot.Pin7.High()
}

func (bot *MyBot) SpinLeft() {
	fmt.Println("Spinning Left!")
	bot.Pin15.Low()
	bot.Pin11.High()
	bot.Pin13.High()
	bot.Pin7.Low()
}

func (bot *MyBot) Close() {
	rpio.Close()
}

func (c *Camera) Hflip(b bool) {
	c.horizontalFlip = b
}

func (c *Camera) Vflip(b bool) {
	c.verticalFlip = b
}

func (bot *MyBot) Capture() (string, error) {
	args := make([]string, 0)
	if bot.CStill.horizontalFlip {
		args = append(args, HFLIP)
	}
	if bot.CStill.verticalFlip {
		args = append(args, VFLIP)
	}
	args = append(args, OUTFLAG)
	fileName := time.Now().Format(TIME_STAMP) + FILE_TYPE
	fullPath := filepath.Join(bot.CStill.savePath, fileName)
	args = append(args, fullPath)
	cmd := exec.Command(STILL, OUTFLAG, fullPath)
	_, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}
	err = cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	cmd.Wait()
	return fullPath, nil
}
