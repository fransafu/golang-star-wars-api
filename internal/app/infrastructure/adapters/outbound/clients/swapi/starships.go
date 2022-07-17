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

type swapiStarshipsClients struct{}

func NewSWAPIStarshipClient() ports.SWAPIStarships {
	return &swapiStarshipsClients{}
}

func (s *swapiStarshipsClients) GetStarships(ctx context.Context) ([]domain.Starship, error) {
	fmt.Println(baseURL + "/starships")
	response, err := http.Get(baseURL + "/starships")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseStarships []domain.Starship
	err = json.Unmarshal(responseData, &responseStarships)
	if err != nil {
		return nil, errors.New("SWAPI CLIENT Starships, GetStarships, ERROR - unmarshal error")
	}

	fmt.Println("results found: %d", len(responseStarships))

	return responseStarships, nil
}

func (s *swapiStarshipsClients) GetStarshipByID(ctx context.Context, req domain.StarshipSearch) (domain.Starship, error) {
	fmt.Println(baseURL + fmt.Sprintf("/starships/%d", req.StarshipID))
	response, err := http.Get(baseURL + fmt.Sprintf("/starships/%d", req.StarshipID))
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseStarship domain.Starship
	err = json.Unmarshal(responseData, &responseStarship)
	if err != nil {
		return responseStarship, errors.New("SWAPI CLIENT Starships, GetStarshipByID, ERROR - unmarshal error")
	}

	return responseStarship, nil
}

// WIP
// func (s *swapiStarshipsClients) GetSchema(ctx context.Context) ()
