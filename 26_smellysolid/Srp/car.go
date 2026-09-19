package srp

import (
	"encoding/json"
	"os"
)

// Location is a simple lat/lng pair.
type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// Car mixes domain behaviour (mileage/travel) with a persistence concern
// (Save) -- two different reasons to change bundled into one struct.
// SRP violation.
type Car struct {
	mileage  int
	location Location
}

// NewCar returns a Car at zero mileage and the zero Location.
func NewCar() *Car {
	return &Car{}
}

// CurrentMileage returns the car's current mileage.
func (c *Car) CurrentMileage() int {
	return c.mileage
}

// TravelTo moves the car to location, incrementing its mileage.
func (c *Car) TravelTo(location Location) {
	c.location = location
	c.mileage++
}

// Save persists the car's state to disk.
func (c *Car) Save() error {
	type carRecord struct {
		Mileage  int      `json:"mileage"`
		Location Location `json:"location"`
	}
	row, err := json.Marshal(carRecord{Mileage: c.mileage, Location: c.location})
	if err != nil {
		return err
	}
	return os.WriteFile("/tmp/car.json", row, 0o644)
}
