package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

// VehicleData represents a single data point for a teleoperated vehicle.
type VehicleData struct {
	VehicleID string
	Timestamp time.Time
	Location  string
	Speed     float64
}

const filePath = "vehicle_data.csv"

func saveDataToFile(data VehicleData, writer *csv.Writer) error {
	record := []string{
		data.VehicleID,
		data.Timestamp.Format(time.RFC3339),
		data.Location,
		strconv.FormatFloat(data.Speed, 'f', 2, 64),
	}

	if err := writer.Write(record); err != nil {
		return err
	}

	writer.Flush()

	return nil
}

func simulateDriving() {
	northLocation := 40.7128
	westLocation := 74.0060
	startLocation := fmt.Sprintf("%f N, %f W", northLocation, westLocation)

	startSpeed := 1.0

	locationIterator := 0.1
	speedIterator := 3.0

	vehicleData := []VehicleData{
		{
			VehicleID: "Vehicle001",
			Timestamp: time.Now(),
			Location:  startLocation,
			Speed:     startSpeed,
		},
	}

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for {
		northLocation += locationIterator
		westLocation += locationIterator
		location := fmt.Sprintf("%f N, %f W", northLocation, westLocation)
		speed := vehicleData[len(vehicleData)-1].Speed + speedIterator

		vehicleData = append(vehicleData, VehicleData{
			VehicleID: "Vehicle001",
			Timestamp: time.Now(),
			Location:  location,
			Speed:     speed,
		})

		if err := saveDataToFile(vehicleData[len(vehicleData)-1], writer); err != nil {
			fmt.Println("Error saving data:", err)
			continue
		}

		fmt.Println("Data saved successfully.")

		locationIterator += 0.1
		speedIterator += 3.0

		time.Sleep(time.Second * 1)
	}
}

func main() {
	simulateDriving()
}
