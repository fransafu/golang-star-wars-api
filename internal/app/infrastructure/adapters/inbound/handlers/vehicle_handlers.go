package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	domain "github.com/fransafu/golang-dragonflydb-example/internal/app/core/domain"
	"github.com/fransafu/golang-dragonflydb-example/internal/app/core/ports"
	"github.com/go-chi/chi/v5"
)

type vehiclesHandlers struct {
	service ports.VehiclesService
}

func NewVehicleHandler(service ports.VehiclesService) ports.Vehicles {
	return &vehiclesHandlers{
		service: service,
	}
}

func (c *vehiclesHandlers) GetVehicles(w http.ResponseWriter, r *http.Request) {
	vehicles, err := c.service.GetVehicles(r.Context())
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(vehicles)
}

func (c *vehiclesHandlers) GetVehicleByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	requestQuery := &domain.VehicleSearch{
		VehicleID: int64(id),
	}

	fmt.Print("%v", requestQuery)

	vehicle, err := c.service.GetVehicleByID(r.Context(), *requestQuery)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(vehicle)
}
