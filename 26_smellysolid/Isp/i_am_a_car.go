package isp

// Location is a simple lat/lng pair.
type Location struct {
	Lat float64
	Lng float64
}

// IAmACar bundles gasoline and electric refuelling -- mutually exclusive
// capabilities -- into one fat interface; no single car honestly supports
// both. ISP violation.
type IAmACar interface {
	GoTo(location Location)
	RefillGasoline(gallons float64) error
	RefillElectricity(kiloWatts float64)
	CurrentMileage() int
}
