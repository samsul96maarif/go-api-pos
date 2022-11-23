package goapiapp

import (
	"samsul96maarif/github.com/go-api-app/lib"
	"samsul96maarif/github.com/go-api-app/repository"
	"samsul96maarif/github.com/go-api-app/usecase"
)

type BE struct {
	Usecase *usecase.Usecase
}

func NewBe(db *lib.Database) (api BE) {
	repo := repository.NewRepository(db)
	usecase := usecase.NewUsecase(&repo)
	api.Usecase = &usecase
	return api
}
