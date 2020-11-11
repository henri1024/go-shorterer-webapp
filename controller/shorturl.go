package controller

import (
	"go-shorterer/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type Query struct {
// 	apikey string `form:"apikey"`
// }

func (suc *MainController) CreateNewShorterer(c *gin.Context) {
	shortlink := &model.ShortLink{}
	if err := c.ShouldBindJSON(shortlink); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid request",
		})
		return
	}

	// query := &Query{}

	// if err := c.ShouldBindQuery(query); err != nil || query.apikey == "" {
	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"msg": "invalid api key",
	// 	})
	// 	return
	// }
	var apikey string

	if apikey = c.Query("apikey"); apikey == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "invalid api key",
		})
		return
	}

	if pass := suc.repo.CheckAPIKey(apikey); !pass {
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
