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

type swapiPeopleClients struct{}

func NewSWAPIPeopleClient() ports.SWAPIPeople {
	return &swapiPeopleClients{}
}

func (s *swapiPeopleClients) GetPeople(ctx context.Context) ([]domain.Person, error) {
	fmt.Println(baseURL + "/films")
	response, err := http.Get(baseURL + "/films")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responsePeople []domain.Person
	err = json.Unmarshal(responseData, &responsePeople)
	if err != nil {
		return nil, errors.New("SWAPI CLIENT People, GetPeople, ERROR - unmarshal error")
	}

	fmt.Println("results found: %d", len(responsePeople))

	return responsePeople, nil
}

func (s *swapiPeopleClients) GetPersonByID(ctx context.Context, req domain.PersonSearch) (domain.Person, error) {
	fmt.Println(baseURL + fmt.Sprintf("/people/%d", req.PersonID))
	response, err := http.Get(baseURL + fmt.Sprintf("/people/%d", req.PersonID))
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responsePerson domain.Person
	err = json.Unmarshal(responseData, &responsePerson)
	if err != nil {
		return responsePerson, errors.New("SWAPI CLIENT FILMS, GetPersonByID, ERROR - unmarshal error")
	}

	return responsePerson, nil
}

// WIP
// func (s *swapiPeopleClients) GetSchema(ctx context.Context) ()
