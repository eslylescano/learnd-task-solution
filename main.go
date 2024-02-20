package main

type Building struct {
	Name     string
	Customer string
	SerialID string
}

type PowerMeterData struct {
	SerialID          string
	ConsumptionPerDay float64
}
