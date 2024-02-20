package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Building struct {
	Name     string
	Customer string
	SerialID string
}

type PowerMeterData struct {
	SerialID          string
	ConsumptionPerDay float64
}

var powerMeters = map[string]Building{
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

func findPowerMetersByCustomer(customer string) []Building {
	var powerMetersForCustomer []Building

	for _, meter := range powerMeters {
		if meter.Customer == customer {
			powerMetersForCustomer = append(powerMetersForCustomer, meter)
		}
	}

	return powerMetersForCustomer
}

func getPowerMetersByCustomer(c *gin.Context) {
	customer := c.Param("customer")
	powerMetersForCustomer := findPowerMetersByCustomer(customer)
	c.JSON(http.StatusOK, gin.H{"powerMeters": powerMetersForCustomer})
}

func getPowerMeterData(c *gin.Context) {
	serialID := c.Param("serialID")
	meter, ok := findPowerMeterData(serialID)

	if ok {
		c.JSON(http.StatusOK, gin.H{"powerMeterData": meter})
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Power meter not found"})

}

func main() {
	router := gin.Default()

	router.GET("/power-meters/:customer", getPowerMetersByCustomer)
	router.GET("/power-meter/:serialID", getPowerMeterData)
	err := router.Run(":8080")
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
