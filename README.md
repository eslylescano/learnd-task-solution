# Power Meter API

This repository contains a simple Go application that implements a Power Meter API using the Gin framework. The API allows you to retrieve information about power meters in different buildings.

## How to Test the Code

### Prerequisites
1. Install [Go](https://golang.org/doc/install) on your machine.
2. Ensure you have a stable internet connection.

### Install Gin Framework
```bash
go get -u github.com/gin-gonic/gin
```

### Clone the Repository
git clone https://github.com/eslylescano/learnd-task-solution.git
cd learnd-task-solution

### Run the Application
```bash
go run main.go
```
The application will start running on http://localhost:8080.

### Test Endpoints

1. Get Power Meters by Customer
```bash
curl http://localhost:8080/power-meters/Aquaflow
```
Response should be:
{
"powerMeters": [
{
"Name": "Treatment Plant A",
"Customer": "Aquaflow",
"SerialID": "1111-1111-1111"
},
{
"Name": "Treatment Plant B",
"Customer": "Aquaflow",
"SerialID": "1111-1111-2222"
}
]
}

2. Get Power Meter Data
```bash
curl http://localhost:8080/power-meter/1111-1111-1111
```
Response should be:{
"powerMeterData": {
"SerialID": "1111-1111-1111",
"ConsumptionPerDay": 20
}
}