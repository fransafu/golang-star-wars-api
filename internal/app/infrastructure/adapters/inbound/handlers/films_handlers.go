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

type filmHandlers struct {
	service ports.FilmService
}

func NewFilmHandler(service ports.FilmService) ports.Film {
	return &filmHandlers{
		service: service,
	}
}

func (c *filmHandlers) GetFilms(w http.ResponseWriter, r *http.Request) {
	films, err := c.service.GetFilms(r.Context())
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(films)
}

func (c *filmHandlers) GetFilmByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	requestQuery := &domain.FilmSearch{
		FilmID: int64(id),
	}

	fmt.Print("%v", requestQuery)

	film, err := c.service.GetFilmByID(r.Context(), *requestQuery)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(film)
}
