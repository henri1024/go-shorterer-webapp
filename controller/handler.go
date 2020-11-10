package controller

import "go-shorterer/repository"

type MainController struct {
	repo repository.DB
}

func NewMainController(repository repository.DB) *MainController {
	return &MainController{
		repo: repository,
	}
}
