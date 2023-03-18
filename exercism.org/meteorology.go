package main

import (
	"fmt"
	"log"
)

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

func (tu TemperatureUnit) String() string {
	switch tu {
	case Celsius:
		return "°C"
	case Fahrenheit:
		return "°F"
	}

	return ""
}

// Add a String method to the TemperatureUnit type

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

func (t Temperature) String() string {
	return fmt.Sprintf("%d %s", t.degree, t.unit)
}

// Add a String method to the Temperature type

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

func (su SpeedUnit) String() string {
	switch su {
	case KmPerHour:
		return "km/h"
	case MilesPerHour:
		return "mph"
	}

	return ""
}

// Add a String method to SpeedUnit

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

func (s Speed) String() string {
	return fmt.Sprintf("%d %s", s.magnitude, s.unit)
}

// Add a String method to Speed

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
func (m MeteorologyData) String() string {
	return fmt.Sprintf("%s: %s, Wind %s at %s, %d%% Humidity", m.location, m.temperature, m.windDirection, m.windSpeed, m.humidity)
}
func main() {
	celsiusUnit := Celsius
	fahrenheitUnit := Fahrenheit

	log.Println(celsiusUnit.String())

	log.Println(fahrenheitUnit.String())
	// => °F
	log.Println(fmt.Sprint(celsiusUnit))
	// Output: °C

	celsiusTemp := Temperature{
		degree: 21,
		unit:   Celsius,
	}
	log.Println(celsiusTemp.String())
	// => 21 °C
	log.Println(fmt.Sprint(celsiusTemp))
	// Output: 21 °C

	fahrenheitTemp := Temperature{
		degree: 75,
		unit:   Fahrenheit,
	}
	log.Println(fahrenheitTemp.String())
	// => 75 °F
	log.Println(fmt.Sprint(fahrenheitTemp))

	mphUnit := MilesPerHour
	log.Println(mphUnit.String())
	// => mph
	log.Println(fmt.Sprint(mphUnit))
	// Output: mph

	kmhUnit := KmPerHour
	log.Println(kmhUnit.String())
	// => km/h
	log.Println(fmt.Sprint(kmhUnit))
	// Output: km/h

	sfData := MeteorologyData{
		location: "San Francisco",
		temperature: Temperature{
			degree: 57,
			unit:   Fahrenheit,
		},
		windDirection: "NW",
		windSpeed: Speed{
			magnitude: 19,
			unit:      MilesPerHour,
		},
		humidity: 60,
	}

	log.Println(sfData.String())
	// => San Francisco: 57 °F, Wind NW at 19 mph, 60% Humidity
	log.Println(fmt.Sprint(sfData))
	// Output: San Francisco: 57 °F, Wind NW at 19 mph, 60% Humidity
}
