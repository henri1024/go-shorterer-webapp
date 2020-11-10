package controller

import (
	"fmt"
	"go-shorterer/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Query struct {
	apikey string `form:"apikey" xml:"apikey" json:"apikey" binding:"required"`
}

func (suc *MainController) CreateNewShorterer(c *gin.Context) {
	shortlink := &model.ShortLink{}
	if err := c.ShouldBindJSON(shortlink); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid request",
		})
		return
	}

	query := &Query{}

	if err := c.ShouldBind(query); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": fmt.Sprint("no api key, error : \v", err),
		})
		return
	}

	log.Printf("received api key : %v\n", query.apikey)

	if pass := suc.repo.CheckAPIKey(query.apikey); !pass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "invalid api key",
		})
		return
	}

	flag := shortlink.SourceKey == ""

	if err := shortlink.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	if err := suc.repo.SaveShortlink(shortlink, flag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"msg": shortlink.ToPublic(),
	})
}

func (suc *MainController) AccessShorterer(c *gin.Context) {
	var key string
	if err := c.ShouldBindUri(&key); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid key",
		})
	}

	var (
		k   string
		err error
	)
	log.Print("\\s\n", key)
	if k, err = suc.repo.GetDestination(key); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid request",
		})
		return
	}

	log.Print("\\s\n", k)

	c.Redirect(http.StatusTemporaryRedirect, k)
	return
}
