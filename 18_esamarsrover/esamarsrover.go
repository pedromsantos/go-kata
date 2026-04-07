package esamarsrover

// Rover - implement this following TDD
var Rover = "Mars plateau"

// RoverInterface for the rover implementation
type RoverInterface interface {
	Execute()
}

// Radio interface for communication
type Radio interface {
	Send(message string)
	Receive() string
}
