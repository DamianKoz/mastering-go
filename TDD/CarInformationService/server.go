package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

type VehicleStatusHandler interface {
	GetVehicleStatus(vehicleID uint) (VehicleStatus, error)
	PostVehicleStatus(vehicleID uint, status VehicleStatus) (VehicleStatus, error)
}

type VehicleServer struct {
	handler VehicleStatusHandler
}

type JSONResponse struct {
	Error         bool          `json:"error"`
	Message       string        `json:"message"`
	VehicleStatus VehicleStatus `json:"vehicleStatus"`
}

type VehicleStatus struct {
	VehicleId    uint         `json:"vehicleId"`
	FuelLevel    int          `json:"fuelLevel"`
	BatteryLevel int          `json:"batteryLevel"`
	EngineStatus string       `json:"engineStatus"`
	SensorStatus SensorStatus `json:"sensorStatus"`
}

type SensorStatus struct {
	FrontCamera string `json:"frontCamera"`
	RearCamera  string `json:"rearCamera"`
	Radar       string `json:"radar"`
	Lidar       string `json:"lidar"`
}

func (vs *VehicleServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	carId, err := validateCarIdFromRequestPath(r.URL.Path)
	if err != nil {
		writeJson(w, 400, JSONResponse{
			Error:         true,
			Message:       "invalid input",
			VehicleStatus: VehicleStatus{},
		})
		return
	}

	switch r.Method {
	case http.MethodGet:
		vehicleStatus, err := vs.handler.GetVehicleStatus(vehicleID)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, JSONResponse{
				Error:   true,
				Message: err.Error(),
			})
			return
		}
		writeJSON(w, http.StatusOK, JSONResponse{
			Error:         false,
			VehicleStatus: vehicleStatus,
		})	case http.MethodPost:
		newStatus := VehicleStatus{}
		vs.handler.PostVehicleStatus(uint(carId), newStatus)
	}

}

func validateCarIdFromRequestPath(path string) (int, error) {
	carId, err := strconv.Atoi(strings.TrimPrefix(path, "/cars/"))
	if err != nil || carId < 1 {
		return 0, errors.New("invalid input")
	}
	return carId, nil
}

func (vs *VehicleServer) GetVehicleStatus(w http.ResponseWriter, r *http.Request) {

	// Here should be the data fetch
	vehicleStatus := VehicleStatus{
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
	}

	jsonResponse := JSONResponse{
		Error:         false,
		Message:       "",
		VehicleStatus: vehicleStatus,
	}

	writeJson(w, 202, jsonResponse)
}

func (vs *VehicleServer) PostVehicleStatus(w http.ResponseWriter, r *http.Request) {
	_, err := validateCarIdFromRequestPath(r.URL.Path)
	if err != nil {
		writeJson(w, 400, JSONResponse{
			Error:         true,
			Message:       "invalid input",
			VehicleStatus: VehicleStatus{},
		})
		return
	}

	// logic for posting new data
	jsonResponse := JSONResponse{
		Error:   false,
		Message: "successfully created",
		VehicleStatus: VehicleStatus{
			VehicleId:    2,
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
	}

	writeJson(w, http.StatusAccepted, jsonResponse)
}

func writeJson(w http.ResponseWriter, status int, data JSONResponse) {
	out, err := json.Marshal(data)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)

	_, err = w.Write(out)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
