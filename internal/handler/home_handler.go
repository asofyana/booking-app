package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowHomePage(c *gin.Context) {

	c.HTML(http.StatusOK, "home.html", gin.H{
		"title": "Home Page"})

}
