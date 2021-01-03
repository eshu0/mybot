package textbot

import (
	"fmt"

	"path/filepath"
	"time"
)

type TextBot struct {
	horizontalFlip bool
	verticalFlip   bool
	savePath       string
}

func NewTextBot(folder string) *TextBot {
	tb := &TextBot{}
	tb.savePath = folder
	return tb
}

func (bot *TextBot) Forwards() {
	fmt.Println("Forwards!")
}

func (bot *TextBot) Backwards() {
	fmt.Println("Backwards!")
}

func (bot *TextBot) Stop() {
	fmt.Println("Stop!")
}

func (bot *TextBot) SpinRight() {
	fmt.Println("Spinning Right!")
}

func (bot *TextBot) SpinLeft() {
	fmt.Println("Spinning Left!")
}

func (bot *TextBot) Close() {
	fmt.Println("CLosing!")
}

func (bot *TextBot) Hflip(b bool) {
	fmt.Println("horizontalFlip is %b", b)
}

func (bot *TextBot) Vflip(b bool) {
	fmt.Println("verticalFlip is %b", b)
}

func (bot *TextBot) Capture() (string, error) {
	fmt.Println("Capturing Photo!")

	args := make([]string, 0)

	if bot.verticalFlip {
		fmt.Println("Flipping V")
	}

	if bot.horizontalFlip {
		fmt.Println("Flipping H")
	}

	FILE_TYPE := ".jpg"
	TIME_STAMP := "2006-01-02_15:04:05"

	fileName := time.Now().Format(TIME_STAMP) + FILE_TYPE
	fullPath := filepath.Join(bot.savePath, fileName)
	fmt.Println("fullPath:")
	fmt.Println(fullPath)
	return fullPath, nil
}
