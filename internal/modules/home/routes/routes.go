package routes

import (
	"github.com/gin-gonic/gin"
	homeCtrl "golang_blog/internal/modules/home/controllers"
)

func Routes(r *gin.Engine) {
	homeController := homeCtrl.NEw()
	r.GET("/", homeController.Index)
}
