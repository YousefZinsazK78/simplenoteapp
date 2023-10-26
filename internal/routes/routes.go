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
		db        = database.NewDatabase(conn)
		userstore = database.NewUserStore(*db)
		hndler    = handler.NewHandler(userstore)

		router = gin.Default()
		// apiV1  = router.Group("/api/v1")
		admin = router.Group("/admin")
	)
	router.Use(cors.Default())

	//todo : login,register => jwt token
	//todo : note crud
	//todo : validation

	//simple note view
	admin.GET("/users", hndler.HandleGetUsers)
	admin.GET("/user/:username", hndler.HandleGetUserByUsername)
	admin.GET("/user/:id", hndler.HandleGetUserById)
	admin.POST("/user", hndler.HandleInsertUser)
	admin.PUT("/user", hndler.HandleUpdateUser)
	admin.DELETE("/user/:id", hndler.HandleDeleteUser)

	return router
}
