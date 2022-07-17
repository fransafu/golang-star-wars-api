package swapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	domain "github.com/fransafu/golang-dragonflydb-example/internal/app/core/domain"
	ports "github.com/fransafu/golang-dragonflydb-example/internal/app/core/ports"
)

type swapiVehiclesClients struct{}

func NewSWAPIVehicleClient() ports.SWAPIVehicles {
	return &swapiVehiclesClients{}
}

func (s *swapiVehiclesClients) GetVehicles(ctx context.Context) (domain.VehicleGetAllResponse, error) {
	fmt.Println(baseURL + "/vehicles")
	response, err := http.Get(baseURL + "/vehicles")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseVehicles domain.VehicleGetAllResponse
	err = json.Unmarshal(responseData, &responseVehicles)
	if err != nil {
		return responseVehicles, errors.New("SWAPI CLIENT Vehicles, GetVehicles, ERROR - unmarshal error")
	}

	fmt.Println("results found: %d", len(responseVehicles.Results))

	return responseVehicles, nil
}

func (s *swapiVehiclesClients) GetVehicleByID(ctx context.Context, req domain.VehicleSearch) (domain.Vehicle, error) {
	fmt.Println(baseURL + fmt.Sprintf("/vehicles/%d", req.VehicleID))
	response, err := http.Get(baseURL + fmt.Sprintf("/vehicles/%d", req.VehicleID))
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseVehicle domain.Vehicle
	err = json.Unmarshal(responseData, &responseVehicle)
	if err != nil {
		return responseVehicle, errors.New("SWAPI CLIENT Vehicle, GetVehicleByID, ERROR - unmarshal error")
	}

	return responseVehicle, nil
}

// WIP
// func (s *swapiVehiclesClients) GetSchema(ctx context.Context) ()
