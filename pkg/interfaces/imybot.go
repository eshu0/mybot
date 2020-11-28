package ibot

// main interface for the SimpleLogger
type IMyBot interface {
	Forwards()
	Backwards()

	Stop()

	SpinRight()
	SpinLeft()

	Close()

	Hflip(b bool)
	Vflip(b bool)
	Capture() (string, error)
}
