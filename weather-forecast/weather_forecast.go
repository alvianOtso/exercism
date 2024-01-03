// Package weather is a package used in this module.
package weather

// CurrentCondition variable used for know current weather.
var CurrentCondition string
// CurrentLocation variable used for know current location of weather.
var CurrentLocation string

// Forecast return current location and condition of weather.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
