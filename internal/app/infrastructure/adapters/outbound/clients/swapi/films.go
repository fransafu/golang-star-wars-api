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

type swapiFilmsClients struct{}

func NewSWAPIFilmsClient() ports.SWAPIFilms {
	return &swapiFilmsClients{}
}

func (s *swapiFilmsClients) GetFilms(ctx context.Context) (domain.FilmGetAllResponse, error) {
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

	var responseFilms domain.FilmGetAllResponse
	err = json.Unmarshal(responseData, &responseFilms)
	if err != nil {
		return responseFilms, errors.New("SWAPI CLIENT FILMS, GetFilms, ERROR - unmarshal error")
	}

	fmt.Println("results found: %d", len(responseFilms.Results))

	return responseFilms, nil
}

func (s *swapiFilmsClients) GetFilmByID(ctx context.Context, req domain.FilmSearch) (domain.Film, error) {
	fmt.Println(baseURL + fmt.Sprintf("/films/%d", req.FilmID))
	response, err := http.Get(baseURL + fmt.Sprintf("/films/%d", req.FilmID))
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseFilm domain.Film
	err = json.Unmarshal(responseData, &responseFilm)
	if err != nil {
		return responseFilm, errors.New("SWAPI CLIENT FILMS, GetFilmByID, ERROR - unmarshal error")
	}

	return responseFilm, nil
}

// WIP
// func (s *swapiFilmsClients) GetSchema(ctx context.Context) ()
