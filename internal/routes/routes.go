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
		apiV1  = router.Group("/api/v1")
	)
	router.Use(cors.Default())
	router.Use(gin.Logger())

	//simple note view
	apiV1.GET("/notes", hndler.HandleGetNotes)
	apiV1.POST("/notes", hndler.HandleGetNotes)
	apiV1.PUT("/notes", hndler.HandleGetNotes)

	return router
}
