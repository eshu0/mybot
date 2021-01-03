package rpibot

import (
	"fmt"

	"os/exec"
	"path/filepath"
	"time"

	ibot "github.com/eshu0/mybot/pkg/interfaces"

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

type PiBot struct {
	ibot.IMyBot
	Pin7   rpio.Pin
	Pin11  rpio.Pin
	Pin13  rpio.Pin
	Pin15  rpio.Pin
	CStill *Camera
}

func NewPiBot(folder string) *PiBot {

	mbot := &PiBot{}
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

	c := &Camera{verticalFlip: true, horizontalFlip: true, savePath: folder}
	mbot.CStill = c
	return mbot

}

func (bot *PiBot) Forwards() {
	fmt.Println("Forwards!")
	bot.Pin15.High()
	bot.Pin11.High()
	bot.Pin13.Low()
	bot.Pin7.Low()
}

func (bot *PiBot) Backwards() {
	fmt.Println("Backwards!")
	bot.Pin15.Low()
	bot.Pin11.Low()
	bot.Pin13.High()
	bot.Pin7.High()
}

func (bot *PiBot) Stop() {
	fmt.Println("Stop!")
	bot.Pin15.Low()
	bot.Pin11.Low()
	bot.Pin13.Low()
	bot.Pin7.Low()
}

func (bot *PiBot) SpinRight() {
	fmt.Println("Spinning Right!")
	bot.Pin15.High()
	bot.Pin11.Low()
	bot.Pin13.Low()
	bot.Pin7.High()
}

func (bot *PiBot) SpinLeft() {
	fmt.Println("Spinning Left!")
	bot.Pin15.Low()
	bot.Pin11.High()
	bot.Pin13.High()
	bot.Pin7.Low()
}

func (bot *PiBot) Close() {
	rpio.Close()
}

func (bot *PiBot) Hflip(b bool) {
	bot.CStill.horizontalFlip = b
}

func (bot *PiBot) Vflip(b bool) {
	bot.CStill.verticalFlip = b
}

func (bot *PiBot) Capture() (string, error) {
	fmt.Println("Capturing Photo!")

	args := make([]string, 0)

	if bot.CStill.verticalFlip {
		args = append(args, VFLIP)
		fmt.Println("Flipping V")
	}

	if bot.CStill.horizontalFlip {
		args = append(args, HFLIP)
		fmt.Println("Flipping H")

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
