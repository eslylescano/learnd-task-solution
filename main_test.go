package main

import "testing"

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
