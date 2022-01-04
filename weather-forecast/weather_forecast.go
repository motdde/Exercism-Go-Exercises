//Package weather is used internally for all weather dealings 👀.
package weather

//CurrentCondition of the city.
var CurrentCondition string

//CurrentLocation lat and long of the city.
var CurrentLocation string

//Forecast returns a mashup of CurrentCondition and CurrentLocation.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
