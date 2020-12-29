package ibot

// main interface for the IMyBot
type IMyBot interface {

	// Move Robot Fowards
	Forwards()

	// Move Robot Backwards
	Backwards()

	// Stop Robots Movements
	Stop()

	// Spin Robot Right
	SpinRight()

	// Spin Robot Left
	SpinLeft()

	// Close Connection
	Close()

	// Set Horizontal Flip of the Camera
	Hflip(b bool)

	// Set Vertical Flip of the Camera
	Vflip(b bool)

	// Returns fie path of camera and error if there was one
	Capture() (string, error)
}
