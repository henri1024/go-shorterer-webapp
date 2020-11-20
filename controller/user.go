package controller

import (
	"go-shorterer/model"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (suc *MainController) CreateUserAPIKEY(c *gin.Context) {

	user := &model.User{}
	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid format",
		})
		return
	}

	err := user.Validate()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	user.APIKEY = uuid.New().String()
	if err = suc.repo.SaveUser(user); err != nil {
		var msg string
		if strings.Contains(err.Error(), "duplicate key value") {
			msg = "Email Already Registered"
		} else {
			msg = "Something went wrong"
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"msg": msg,
		})
		return
	}

	log.Printf("user api key created : %v\n", user)

	// err = suc.mail.Send(user.APIKEY, user.Email)

	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"msg": err.Error(),
	// 	})
	// 	return
	// }

	c.JSON(http.StatusCreated, gin.H{
		"msg": user.APIKEY,
	})

	return
}
