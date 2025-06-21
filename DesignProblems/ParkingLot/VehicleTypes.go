package parkinglot

// like Class in Java
type VehicleType string

// like Enum in Java
const (
	Bike  VehicleType = "Bike"
	Car   VehicleType = "Car"
	Truck VehicleType = "Truck"
)

type Vehicle struct {
	Number string
	Type   VehicleType
}
