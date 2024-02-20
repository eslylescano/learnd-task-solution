package main

import "testing"

type Building struct {
	Name     string
	Customer string
	SerialID string
}

type PowerMeterData struct {
	SerialID          string
	ConsumptionPerDay float64
}

var powerMetters = map[string]Building{
	"1111-1111-1111": {"Treatment Plant A", "Aquaflow", "1111-1111-1111"},
	"1111-1111-2222": {"Treatment Plant B", "Aquaflow", "1111-1111-2222"},
	"1111-1111-3333": {"Student Halls", "Albers Facilities Management", "1111-1111-3333"},
}

var meterConsumption = map[string]PowerMeterData{
	"1111-1111-1111": {"1111-1111-1111", 20.0},
	"1111-1111-2222": {"1111-1111-2222", 30.0},
	"1111-1111-3333": {"1111-1111-3333", 40.0},
}

func TestFindPowerMeterData(t *testing.T) {
	serialID := "1111-1111-1111"
	expectedMeterData := PowerMeterData{SerialID: "1111-1111-1111", ConsumptionPerDay: 20.0}

	meter, ok := findPowerMeterData(serialID)

	if !ok {
		t.Errorf("Expected power meter data for serialID %s, but not found", serialID)
	}

	if meter != expectedMeterData {
		t.Errorf("Expected %v, got %v", expectedMeterData, meter)
	}
}
