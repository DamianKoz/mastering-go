package main

import (
	"errors"
	"log"
	"net/http"
)

type InMemoryVehicleStatusHandler struct {
	vehicles map[uint]VehicleStatus
}

func NewInMemoryVehicleStatusHandler() *InMemoryVehicleStatusHandler {
	return &InMemoryVehicleStatusHandler{
		vehicles: map[uint]VehicleStatus{
			1: {
				VehicleId:    1,
				FuelLevel:    65,
				BatteryLevel: 40,
				EngineStatus: "Normal",
				SensorStatus: SensorStatus{
					FrontCamera: "Operational",
					RearCamera:  "Operational",
					Radar:       "Operational",
					Lidar:       "Operational",
				},
			},
			2: {
				VehicleId:    2,
				FuelLevel:    70,
				BatteryLevel: 40,
				EngineStatus: "Normal",
				SensorStatus: SensorStatus{
					FrontCamera: "Operational",
					RearCamera:  "Operational",
					Radar:       "Operational",
					Lidar:       "Operational",
				},
			},
		},
	}
}

func (h *InMemoryVehicleStatusHandler) GetVehicleStatus(vehicleID uint) (VehicleStatus, error) {
	status, exists := h.vehicles[vehicleID]
	if !exists {
		return VehicleStatus{}, errors.New("vehicle not found")
	}
	return status, nil
}

func (h *InMemoryVehicleStatusHandler) PostVehicleStatus(vehicleID uint, status VehicleStatus) (VehicleStatus, error) {
	h.vehicles[vehicleID] = status
	return status, nil
}

func main() {
	server := &VehicleServer{handler: NewInMemoryVehicleStatusHandler()}
	log.Fatal(http.ListenAndServe(":8000", server))
}
