package controllers

import (
	"github.com/gin-gonic/gin"
	"golang_blog/pkg/html"
	"net/http"
)

type Controller struct {
}

func NEw() *Controller {
	return &Controller{}
}
func (Controller) Index(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
		"title": "Home page",
	})

}
