// Package weather forecast the current weather condition of various cities in Goblinocus.
package weather

// CurrentCondition represents the current condition of Goblinocus.
var CurrentCondition string

// CurrentLocation represents the current location of Goblinocus.
var CurrentLocation string

// Forecast returns the location and current weater condition
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
