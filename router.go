package main

import (
	"shorterer/controller"

	"github.com/gin-gonic/gin"
)

// Router model
type Router struct {
	Router *gin.Engine
}

// NewRouter to create a router struct and return it as pointer
func NewRouter(controller *controller.ShortUrlController) *Router {
	router := &Router{
		gin.Default(),
	}
	router.AddEndPoint(controller)
	return router
}

// AddEndPoint which serve and connect url path to controller
func (r *Router) AddEndPoint(controller *controller.ShortUrlController) {
	r.Router.POST("api/new", controller.CreateNewShorterer)
	r.Router.GET(":key", controller.AccessShorterer)
}
