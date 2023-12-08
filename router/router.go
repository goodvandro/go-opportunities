package router

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default()
	initializeRoutes(router)
	router.Run() // listen and serve on 0.0.0.0:8080
}
