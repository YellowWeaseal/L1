package main

import "fmt"

//Реализация адаптера на примере температуры в Цельсиях и Фаренгейтах

type CelsiusTemperature interface {
	GetTemperatureCelsius() float64
}

type Celsius struct {
	Temperature float64
}

func (c Celsius) GetTemperatureCelsius() float64 {
	return c.Temperature
}

type FahrenheitTemperature interface {
	GetTemperatureFahrenheit() float64
}

type Fahrenheit struct {
	Temperature float64
}

func (f Fahrenheit) GetTemperatureFahrenheit() float64 {
	return f.Temperature
}

type FahrenheitAdapter struct {
	FahrenheitTemperature
}

func (fa FahrenheitAdapter) GetTemperatureCelsius() float64 {
	// Конвертируем температуру из Фаренгейтов в Цельсии
	celsius := (fa.GetTemperatureFahrenheit() - 32) * 5 / 9
	return celsius
}

func main() {
	celsiusTemperature := Celsius{Temperature: 25.0}
	fahrenheitTemperature := Fahrenheit{Temperature: 77.0}

	// Используем адаптер для конвертации температуры из Фаренгейтов в Цельсии
	adapter := FahrenheitAdapter{FahrenheitTemperature: fahrenheitTemperature}
	celsiusTemperatureFromAdapter := adapter.GetTemperatureCelsius()

	fmt.Printf("Температура в Цельсиях: %.2f\n", celsiusTemperature.GetTemperatureCelsius())
	fmt.Printf("Температура в Цельсиях (через адаптер): %.2f\n", celsiusTemperatureFromAdapter)
}
