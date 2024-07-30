package routing

import (
	"fmt"
	"golang_blog/pkg/config"
	"log"
)

func Serve() {
	r := GetRouter()

	configs := config.Get()

	err := r.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port)) // listen and serve on localhost:8000 (for windows "localhost:8080")

	if err != nil {
		log.Fatalln("Error in routing")

		return
	}
}
