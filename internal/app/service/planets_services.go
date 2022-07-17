package service

import (
	"context"

	domain "github.com/fransafu/golang-dragonflydb-example/internal/app/core/domain"
	"github.com/fransafu/golang-dragonflydb-example/internal/app/core/ports"
)

type planetService struct {
	swapiPlanetClient ports.SWAPIPlanets
}

func NewPlanetService(swapiPlanetClient ports.SWAPIPlanets) ports.PlanetService {
	return &planetService{
		swapiPlanetClient: swapiPlanetClient,
	}
}

func (f *planetService) GetPlanets(ctx context.Context) (domain.PlanetGetAllResponse, error) {
	planets, err := f.swapiPlanetClient.GetPlanets(ctx)
	if err != nil {
		return planets, err
	}

	return planets, nil
}

func (f *planetService) GetPlanetByID(ctx context.Context, search domain.PlanetSearch) (domain.Planet, error) {
	planet, err := f.swapiPlanetClient.GetPlanetByID(ctx, search)
	if err != nil {
		return planet, err
	}

	return planet, nil
}
