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

func TestFindPowerMetersByCustomer(t *testing.T) {
	customer := "Aquaflow"
	expectedCount := 2

	powerMetersForCustomer := findPowerMetersByCustomer(customer)

	if len(powerMetersForCustomer) != expectedCount {
		t.Errorf("Expected %d power meters for customer %s, got %d", expectedCount, customer, len(powerMetersForCustomer))
	}
}
