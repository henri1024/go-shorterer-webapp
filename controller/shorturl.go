package controller

import (
	"go-shorterer/model"
	"go-shorterer/repository"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ShortUrlController struct {
	shortUrlRepository repository.DB
}

func NewShortUrlController(repository repository.DB) *ShortUrlController {
	return &ShortUrlController{
		shortUrlRepository: repository,
	}
}

func (suc *ShortUrlController) CreateNewShorterer(c *gin.Context) {
	shortlink := &model.ShortLink{}
	if err := c.ShouldBindJSON(shortlink); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid request",
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

	if err := suc.shortUrlRepository.Save(shortlink, flag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid request",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"msg": shortlink.ToPublic(),
	})
}

func (suc *ShortUrlController) AccessShorterer(c *gin.Context) {
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
	if k, err = suc.shortUrlRepository.GetDestination(key); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid request",
		})
		return
	}

	log.Print("\\s\n", k)

	c.Redirect(http.StatusTemporaryRedirect, k)
	return
}
