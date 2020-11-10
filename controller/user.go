package controller

import (
	"go-shorterer/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (suc *MainController) CreateUserAPIKEY(c *gin.Context) {

	user := &model.User{}
	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid format",
		})
	}

	err := user.Validate()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	user.APIKEY = uuid.New().String()
	err = suc.repo.SaveUser(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	log.Printf("user api key created : %v\n", user)

	c.JSON(http.StatusCreated, gin.H{
		"api_key": user.APIKEY,
	})

	return
}
