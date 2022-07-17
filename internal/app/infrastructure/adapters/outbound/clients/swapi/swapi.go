package swapi

import (
	ports "github.com/fransafu/golang-dragonflydb-example/internal/app/core/ports"
)

const baseURL = "https://swapi.dev/api"

type swapiClients struct{}

func NewSWAPIClient() ports.SWAPI {
	return &swapiClients{}
}
