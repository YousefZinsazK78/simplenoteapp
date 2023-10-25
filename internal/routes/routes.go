package routes

import (
	"notegin/internal/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	var (
		router = gin.Default()
		hndler = handler.NewHandler()
	)
	router.Use(cors.Default())

	//simple note view
	apiV1 := router.Group("/api/v1")

	apiV1.GET("/notes", hndler.HandleGetNotes)

	return router
}
