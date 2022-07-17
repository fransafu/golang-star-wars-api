package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	handlers "github.com/fransafu/golang-dragonflydb-example/internal/app/infrastructure/adapters/inbound/handlers"
	swapi "github.com/fransafu/golang-dragonflydb-example/internal/app/infrastructure/adapters/outbound/clients/swapi"
	service "github.com/fransafu/golang-dragonflydb-example/internal/app/service"
)

func StartApp() error {
	r := chi.NewRouter()

	swapiFilmsClient := swapi.NewSWAPIFilmsClient()
	swapiPlanetsClient := swapi.NewSWAPIPlanetsClient()
	swapiVehicleClient := swapi.NewSWAPIVehicleClient()

	filmService := service.NewFilmService(swapiFilmsClient)
	planetService := service.NewPlanetService(swapiPlanetsClient)
	vehicleService := service.NewVehicleService(swapiVehicleClient)

	filmHandler := handlers.NewFilmHandler(filmService)
	planetHandler := handlers.NewPlanetHandler(planetService)
	vehicleHandler := handlers.NewVehicleHandler(vehicleService)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Route("/films", func(r chi.Router) {
		r.Get("/", filmHandler.GetFilms)
		r.Get("/{id}", filmHandler.GetFilmByID)
	})

	r.Route("/planets", func(r chi.Router) {
		r.Get("/", planetHandler.GetPlanets)
		r.Get("/{id}", planetHandler.GetPlanetByID)
	})

	r.Route("/vehicles", func(r chi.Router) {
		r.Get("/", vehicleHandler.GetVehicles)
		r.Get("/{id}", vehicleHandler.GetVehicleByID)
	})

	return http.ListenAndServe(":3333", r)
}
