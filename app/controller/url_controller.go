package controller

import (
	"net/http"
	"strconv"

	"github.com/0l1v3rr/gourly/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func GetAllUrls(c *gin.Context) {
	res, err := model.URL{}.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

func GetUrlById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please provide a valid ID.",
		})
		return
	}

	res, err := model.URL{}.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

func CreateUrl(c *gin.Context) {
	var url model.URL

	if err := c.ShouldBindBodyWith(&url, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	created, err := model.URL{}.Create(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, created)
}

func DeleteUrl(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please provide a valid ID.",
		})
		return
	}

	err = model.URL{}.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusAccepted)
}

func IncrementUrlClicks(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please provide a valid ID.",
		})
		return
	}

	url, err := model.URL{}.IncrementClicks(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, url)
}

func Redirect(c *gin.Context) {
	hash := c.Param("hash")

	url, err := model.URL{}.GetByHash(hash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Redirect(http.StatusMovedPermanently, url.RedirectTo)
}
