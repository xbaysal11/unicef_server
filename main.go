package main

import (
	"unicef_server/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	routes.AddRoutes(r)
	r.Run("0.0.0.0:9000") // listen and serve on 0.0.0.0:8080
}
