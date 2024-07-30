package routing

import (
	"golang_blog/internal/providers/routes"

	"github.com/gin-gonic/gin"
)

func Init() {
	r = gin.Default()

}

func GetRouter() *gin.Engine {
	return r
}

func RegisterRoutes() {
	routes.RegisterRoutes(GetRouter())

}
