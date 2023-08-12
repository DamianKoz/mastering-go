package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGETCars(t *testing.T) {
	server := &VehicleServer{handler: NewInMemoryVehicleStatusHandler()}
	t.Run("returns status of one car", func(t *testing.T) {
		request := newGetCarStatusRequest("1")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		want := JSONResponse{
			Error:   false,
			Message: "",
			VehicleStatus: VehicleStatus{
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
		}

		assertResponseBody(t, *response, want)
		assertResponseStatusCode(t, response.Code, http.StatusAccepted)

	})

	t.Run("returns error due to unknown id", func(t *testing.T) {
		request := newGetCarStatusRequest("-1")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		want := JSONResponse{
			Error:         true,
			Message:       "invalid input",
			VehicleStatus: VehicleStatus{},
		}

		assertResponseBody(t, *response, want)
		assertResponseStatusCode(t, response.Code, http.StatusBadRequest)

	})
}

var POST_REQUEST_BODY = VehicleStatus{
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
}

func TestPOSTCars(t *testing.T) {

	server := &VehicleServer{handler: NewInMemoryVehicleStatusHandler()}

	t.Run("returns accepted on POST", func(t *testing.T) {
		request := newPostCarStatusRequest("2")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		want := JSONResponse{
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

		assertResponseBody(t, *response, want)
		assertResponseStatusCode(t, response.Code, http.StatusAccepted)
	})

}

func newGetCarStatusRequest(id string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/cars/%s", id), nil)
	return request
}

func newPostCarStatusRequest(id string) *http.Request {
	body, err := json.Marshal(POST_REQUEST_BODY)
	if err != nil {
		log.Println("Error encoding body: ", body)
	}

	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/cars/%s", id), bytes.NewBuffer(body))
	return request
}

func assertResponseBody(t *testing.T, response httptest.ResponseRecorder, want JSONResponse) {
	t.Helper()
	var got JSONResponse
	err := json.Unmarshal(response.Body.Bytes(), &got)
	if err != nil {
		t.Fatalf("Could not unmarshal response %s into struct: %v", response.Body.String(), err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func assertResponseStatusCode(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("got StatusCode %d, want StatusCode %d", got, want)
	}
}
