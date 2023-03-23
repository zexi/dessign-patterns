package main

func main() {
	wd := NewWeatherData()
	wd.RegisterObserver(NewLineDisplayer())
	wd.RegisterObserver(NewWrapTextDisplayer())

	wd.SetMeasurements(35.2, 10)
	wd.SetMeasurements(25.2, 9)
}
