package main

import "observer"

type Data struct {
	Temperature float64
	Pressure    float64
}

type WeatherData struct {
	subject observer.Subject[*Data]
	data    *Data
}

func NewWeatherData() *WeatherData {
	return &WeatherData{
		subject: observer.NewSubject[*Data](),
		data:    new(Data),
	}
}

func (wd *WeatherData) RegisterObserver(ob observer.Observer[*Data]) {
	wd.subject.RegisterObserver(ob)
}

func (wd *WeatherData) RemoveObserver(ob observer.Observer[*Data]) {
	wd.subject.RemoveObserver(ob)
}

func (wd *WeatherData) NotifyObservers(data *Data) {
	wd.NotifyObservers(data)
}

func (wd *WeatherData) measurementsChanged() {
	wd.subject.NotifyObservers(wd.data)
}

func (wd *WeatherData) SetMeasurements(temp float64, pressure float64) {
	wd.data = &Data{
		Temperature: temp,
		Pressure:    pressure,
	}
	wd.measurementsChanged()
}
