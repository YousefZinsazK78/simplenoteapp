package routes

import (
	"database/sql"
	"log"
	"notegin/internal/database"
	"notegin/internal/handler"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Init() *gin.Engine {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	var conn, err = sql.Open("postgres", os.Getenv("DBCONNECTION"))
	if err != nil {
		log.Fatal(err)
	}

	var (
		db     = database.NewDatabase(conn)
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
