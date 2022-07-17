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

type swapiPlanetClient struct{}

func NewSWAPIPlanetsClient() ports.SWAPIPlanets {
	return &swapiPlanetClient{}
}

func (s *swapiPlanetClient) GetPlanets(ctx context.Context) (domain.PlanetGetAllResponse, error) {
	fmt.Println(baseURL + "/planets")
	response, err := http.Get(baseURL + "/planets")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responsePlanets domain.PlanetGetAllResponse
	err = json.Unmarshal(responseData, &responsePlanets)
	if err != nil {
		return responsePlanets, errors.New("SWAPI CLIENT planets, GetPlanets, ERROR - unmarshal error")
	}

	fmt.Println("results found: %d", len(responsePlanets.Results))

	return responsePlanets, nil
}

func (s *swapiPlanetClient) GetPlanetByID(ctx context.Context, req domain.PlanetSearch) (domain.Planet, error) {
	fmt.Println(baseURL + fmt.Sprintf("/planets/%d", req.PlanetID))
	response, err := http.Get(baseURL + fmt.Sprintf("/planets/%d", req.PlanetID))
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responsePlanet domain.Planet
	err = json.Unmarshal(responseData, &responsePlanet)
	if err != nil {
		return responsePlanet, errors.New("SWAPI CLIENT Planet, GetPlanetByID, ERROR - unmarshal error")
	}

	return responsePlanet, nil
}

// WIP
// func (s *swapiPlanetClient) GetSchema(ctx context.Context) ()
