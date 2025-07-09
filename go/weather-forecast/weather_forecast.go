/*
Package weather is providing functionality to forecast the
weather for Goblinocus.
*/
package weather

// CurrentCondition contains the weather conditions for the CurrentLocation.
var CurrentCondition string

// CurrentLocation conatins the location for the weather forecast.
var CurrentLocation string

/*
Forecast is calculating the current waeather conditions based
on the provided city and condition.
*/
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
