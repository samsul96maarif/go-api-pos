package usecase

import "samsul96maarif/github.com/go-api-app/repository"

type Usecase struct {
	repo *repository.Repository
}

func NewUsecase(repo *repository.Repository) Usecase {
	return Usecase{repo: repo}
}
