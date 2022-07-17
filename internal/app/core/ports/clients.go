package ports

import (
	"context"
	"net/http"

	domain "github.com/fransafu/golang-dragonflydb-example/internal/app/core/domain"
)

//go:generate mockgen -package=mock -source=./clients.go -destination=./mock/clients_mock.go
type Clients interface {
	Film
	People
	Vehicles
	Planet
}

type Film interface {
	GetFilms(w http.ResponseWriter, r *http.Request)
	GetFilmByID(w http.ResponseWriter, r *http.Request)
}

type Planet interface {
	GetPlanets(w http.ResponseWriter, r *http.Request)
	GetPlanetByID(w http.ResponseWriter, r *http.Request)
}

type Vehicles interface {
	GetVehicles(w http.ResponseWriter, r *http.Request)
	GetVehicleByID(w http.ResponseWriter, r *http.Request)
}

type People interface {
	GetFilm(ctx context.Context, filmID uint) (*domain.Film, error)
}

/*
 * Outgoing
 */
type SWAPI interface {
}

type SWAPIFilms interface {
	GetFilms(context.Context) (domain.FilmGetAllResponse, error)
	GetFilmByID(context.Context, domain.FilmSearch) (domain.Film, error)
}

type SWAPIPeople interface {
	GetPeople(context.Context) ([]domain.Person, error)
	GetPersonByID(context.Context, domain.PersonSearch) (domain.Person, error)
}

type SWAPIStarships interface {
	GetStarships(context.Context) ([]domain.Starship, error)
	GetStarshipByID(context.Context, domain.StarshipSearch) (domain.Starship, error)
}

type SWAPIVehicles interface {
	GetVehicles(context.Context) (domain.VehicleGetAllResponse, error)
	GetVehicleByID(context.Context, domain.VehicleSearch) (domain.Vehicle, error)
}

type SWAPISpecies interface {
	GetSpecies(context.Context) ([]domain.Species, error)
	GetSpecieByID(context.Context, domain.SpecieSearch) (domain.Species, error)
}

type SWAPIPlanets interface {
	GetPlanets(context.Context) (domain.PlanetGetAllResponse, error)
	GetPlanetByID(context.Context, domain.PlanetSearch) (domain.Planet, error)
}
