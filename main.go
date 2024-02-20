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

func findPowerMeterData(serialID string) (PowerMeterData, bool) {
	meter, ok := meterConsumption[serialID]
	return meter, ok
}
