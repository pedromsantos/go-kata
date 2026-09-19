package isp

import (
	"errors"
	"fmt"
)

// ElectricCar is the forced implementer: it errors on the gasoline method
// it can't honestly support.
type ElectricCar struct {
	mileage          int
	batteryKiloWatts float64
}

// GoTo drives the car to location, incrementing mileage.
func (c *ElectricCar) GoTo(location Location) {
	c.mileage++
	fmt.Printf("Driving to %v, %v\n", location.Lat, location.Lng)
}

// RefillGasoline always fails: electric cars don't take gasoline.
func (c *ElectricCar) RefillGasoline(_ float64) error {
	return errors.New("electric cars don't take gasoline")
}

// RefillElectricity adds kiloWatts to the battery.
func (c *ElectricCar) RefillElectricity(kiloWatts float64) {
	c.batteryKiloWatts += kiloWatts
}

// CurrentMileage returns the car's current mileage.
func (c *ElectricCar) CurrentMileage() int {
	return c.mileage
}
