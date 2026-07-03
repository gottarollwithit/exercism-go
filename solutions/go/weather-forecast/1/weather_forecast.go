// Package weather gives you the necessary tools to forecast weather conditions.
package weather

var (
	//CurrentCondition stores the current weather condition.
	CurrentCondition string
	//CurrentLocation stores the city of the forecast.
	CurrentLocation string
)

// Forecast returns a string describing current condition of the requested city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
