package handlers

import "github.com/fransafu/golang-dragonflydb-example/internal/app/core/ports"

type Handlers struct {
	service ports.Service
}

func NewHTTPServer(service ports.Service) Handlers {
	return Handlers{
		service: service,
	}
}
