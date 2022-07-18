// Package weather: weather information -.-.
package weather

// CurrentCondition: some var 1.
var CurrentCondition string

// CurrentLocation:  var 2.
var CurrentLocation string

// Forecast: predict weather?.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
