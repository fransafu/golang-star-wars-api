package service

import (
	"context"

	domain "github.com/fransafu/golang-dragonflydb-example/internal/app/core/domain"
	"github.com/fransafu/golang-dragonflydb-example/internal/app/core/ports"
)

type vehicleService struct {
	swapiVehicleClient ports.SWAPIVehicles
}

func NewVehicleService(swapiVehicleClient ports.SWAPIVehicles) ports.VehiclesService {
	return &vehicleService{
		swapiVehicleClient: swapiVehicleClient,
	}
}

func (f *vehicleService) GetVehicles(ctx context.Context) (domain.VehicleGetAllResponse, error) {
	vehicles, err := f.swapiVehicleClient.GetVehicles(ctx)
	if err != nil {
		return vehicles, err
	}

	return vehicles, nil
}

func (f *vehicleService) GetVehicleByID(ctx context.Context, search domain.VehicleSearch) (domain.Vehicle, error) {
	vehicle, err := f.swapiVehicleClient.GetVehicleByID(ctx, search)
	if err != nil {
		return vehicle, err
	}

	return vehicle, nil
}
