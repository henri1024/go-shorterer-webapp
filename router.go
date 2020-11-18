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
func NewRouter(controller *controller.MainController) *Router {
	router := &Router{
		gin.Default(),
	}
	router.Router.Use(CORSMiddleware())
	router.AddEndPoint(controller)
	return router
}

// AddEndPoint which serve and connect url path to controller
func (r *Router) AddEndPoint(controller *controller.MainController) {
	r.Router.POST("api/new", controller.CreateNewShorterer)
	r.Router.POST("api/newkey", controller.CreateUserAPIKEY)
	r.Router.GET(":key", controller.AccessShorterer)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
