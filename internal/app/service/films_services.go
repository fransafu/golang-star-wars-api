package service

import (
	"context"
	"fmt"

	domain "github.com/fransafu/golang-dragonflydb-example/internal/app/core/domain"
	"github.com/fransafu/golang-dragonflydb-example/internal/app/core/ports"
)

type filmService struct {
	swapiFilmsClient ports.SWAPIFilms
}

func NewFilmService(swapiFilmsClient ports.SWAPIFilms) ports.FilmService {
	return &filmService{
		swapiFilmsClient: swapiFilmsClient,
	}
}

func (f *filmService) GetFilms(ctx context.Context) (domain.FilmGetAllResponse, error) {
	films, err := f.swapiFilmsClient.GetFilms(ctx)
	if err != nil {
		fmt.Print(err.Error())
	}

	return films, nil
}

func (f *filmService) GetFilmByID(ctx context.Context, search domain.FilmSearch) (domain.Film, error) {
	film, err := f.swapiFilmsClient.GetFilmByID(ctx, search)
	if err != nil {
		fmt.Print(err.Error())
		return film, err
	}

	return film, nil
}
