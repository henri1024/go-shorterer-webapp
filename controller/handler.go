package controller

import (
	"go-shorterer/app"
	"go-shorterer/repository"
)

type MainController struct {
	repo repository.DB
	mail *app.EmailWidget
}

func NewMainController(repository repository.DB, mailWidget *app.EmailWidget) *MainController {
	return &MainController{
		repo: repository,
		mail: mailWidget,
	}
}
