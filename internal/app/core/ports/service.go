package ports

import (
	"context"

	domain "github.com/fransafu/golang-dragonflydb-example/internal/app/core/domain"
)

//go:generate mockgen -package=mock -source=./service.go -destination=./mock/service_mock.go
type Service interface{}

type FilmService interface {
	GetFilms(context.Context) (domain.FilmGetAllResponse, error)
	GetFilmByID(context.Context, domain.FilmSearch) (domain.Film, error)
}

type PlanetService interface {
	GetPlanets(context.Context) (domain.PlanetGetAllResponse, error)
	GetPlanetByID(context.Context, domain.PlanetSearch) (domain.Planet, error)
}

type PeopleService interface {
	GetPeople(context.Context) ([]domain.Person, error)
	GetPersonByID(context.Context, domain.PersonSearch) (domain.Person, error)
}

type StarshipService interface {
	GetStarships(context.Context) ([]domain.Starship, error)
	GetStarshipByID(context.Context, domain.StarshipSearch) (domain.Starship, error)
}

type VehiclesService interface {
	GetVehicles(context.Context) (domain.VehicleGetAllResponse, error)
	GetVehicleByID(context.Context, domain.VehicleSearch) (domain.Vehicle, error)
}

type SpeciesService interface {
	GetSpecies(context.Context) ([]domain.Species, error)
	GetSpecieByID(context.Context, domain.SpecieSearch) (domain.Species, error)
}
