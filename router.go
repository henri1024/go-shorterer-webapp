package main

import (
	"go-shorterer/controller"

	"github.com/gin-gonic/gin"
)

// Router model
type Router struct {
	Router *gin.Engine
}

// NewRouter to create a router struct and return it as pointer
func NewRouter(controller *controller.MainController, corsMid gin.HandlerFunc) *Router {
	router := &Router{
		gin.Default(),
	}
	router.Router.Use(corsMid)
	router.AddEndPoint(controller)
	return router
}

// AddEndPoint which serve and connect url path to controller
func (r *Router) AddEndPoint(controller *controller.MainController) {
	r.Router.POST("api/new", controller.CreateNewShorterer)
	r.Router.POST("api/newkey", controller.CreateUserAPIKEY)
	r.Router.GET(":key", controller.AccessShorterer)
}
