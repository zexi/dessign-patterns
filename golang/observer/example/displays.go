package main

import (
	"log"
	"observer"
)

type LineDisplayer struct{}

func NewLineDisplayer() observer.Observer[*Data] {
	return &LineDisplayer{}
}

func (ld *LineDisplayer) GetId() string {
	return "line_displayer"
}

func (ld *LineDisplayer) OnChange(t *Data) {
	log.Printf("weather data: {temperature: %f, pressure: %f}", t.Temperature, t.Pressure)
}

type WrapTextDisplayer struct{}

func NewWrapTextDisplayer() observer.Observer[*Data] {
	return &WrapTextDisplayer{}
}

func (ld *WrapTextDisplayer) GetId() string {
	return "wrap_text_displayer"
}

func (ld *WrapTextDisplayer) OnChange(t *Data) {
	log.Printf("=====================")
	log.Printf("{temperature: %f,\npressure: %f}", t.Temperature, t.Pressure)
	log.Printf("=====================")
}
