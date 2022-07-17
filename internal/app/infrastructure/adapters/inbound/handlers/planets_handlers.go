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

type planetsHandlers struct {
	service ports.PlanetService
}

func NewPlanetHandler(service ports.PlanetService) ports.Planet {
	return &planetsHandlers{
		service: service,
	}
}

func (c *planetsHandlers) GetPlanets(w http.ResponseWriter, r *http.Request) {
	planets, err := c.service.GetPlanets(r.Context())
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(planets)
}

func (c *planetsHandlers) GetPlanetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	requestQuery := &domain.PlanetSearch{
		PlanetID: int64(id),
	}

	fmt.Print("%v", requestQuery)

	planet, err := c.service.GetPlanetByID(r.Context(), *requestQuery)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(planet)
}
