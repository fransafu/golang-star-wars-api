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

type swapiSpeciesClients struct{}

func NewSWAPISpeciesClient() ports.SWAPISpecies {
	return &swapiSpeciesClients{}
}

func (s *swapiSpeciesClients) GetSpecies(ctx context.Context) ([]domain.Species, error) {
	fmt.Println(baseURL + "/species")
	response, err := http.Get(baseURL + "/species")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseSpecies []domain.Species
	err = json.Unmarshal(responseData, &responseSpecies)
	if err != nil {
		return nil, errors.New("SWAPI CLIENT Species, GetSpecies, ERROR - unmarshal error")
	}

	fmt.Println("results found: %d", len(responseSpecies))

	return responseSpecies, nil
}

func (s *swapiSpeciesClients) GetSpecieByID(ctx context.Context, req domain.SpecieSearch) (domain.Species, error) {
	fmt.Println(baseURL + fmt.Sprintf("/species/%d", req.SpecieID))
	response, err := http.Get(baseURL + fmt.Sprintf("/species/%d", req.SpecieID))
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseSpecie domain.Species
	err = json.Unmarshal(responseData, &responseSpecie)
	if err != nil {
		return responseSpecie, errors.New("SWAPI CLIENT Specie, GetSpecieByID, ERROR - unmarshal error")
	}

	return responseSpecie, nil
}

// WIP
// func (s *swapiSpeciesClients) GetSchema(ctx context.Context) ()
