package tmplate

import (
	"database/sql"
	"log"
	"net/http"
	"notegin/internal/database"
	"notegin/internal/handler"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Init() *gin.Engine {
	if err := godotenv.Load("./internal/config/.env"); err != nil {
		log.Fatal(err)
	}

	var conn, err = sql.Open("postgres", os.Getenv("DBCONNECTION"))
	if err != nil {
		log.Fatal(err)
	}

	var (
		db        = database.NewDatabase(conn)
		userstore = database.NewUserStore(*db)
		notestore = database.NewNoteStore(*db)
		hndler    = handler.NewHandler(userstore, notestore)

		router = gin.Default()
		// admin  = router.Group("/admin")
		// auth = router.Group("/auth")
		note = router.Group("/note")
	)
	router.Use(cors.Default())
	router.LoadHTMLGlob("./internal/templates/*.tmpl")
	router.Static("/assets", "./internal/assets")
	//testing route uri
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "main_note.tmpl", gin.H{
			"title": "Main Note Page",
		})
	})
	router.GET("/create", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "create.tmpl", nil)
	})

	// auth.GET("/tmplsignin", func(ctx *gin.Context) {
	// 	ctx.HTML(http.StatusOK, "login.tmpl", nil)
	// })
	// auth.POST("/signin", hndler.HandleTmplSignIn)
	// auth.POST("/signup", hndler.HandleSignUp)

	// note.Use(middleware.JwtAuth())
	//note crud
	note.GET("/all", hndler.HandleTmplGetNotes)
	note.GET("/title/:title", hndler.HandleGetNoteTitle)
	note.GET("/:id", hndler.HandleGetNoteByID)
	note.POST("/create", hndler.HandleTmplCreateNote)
	note.PUT("/update", hndler.HandleUpdateNote)
	note.DELETE("/delete/:id", hndler.HandleDeleteNote)

	// admin.Use(middleware.JwtAuth())
	// //user crud
	// admin.GET("/users", hndler.HandleGetUsers)
	// admin.GET("/user/username/:username", hndler.HandleGetUserByUsername)
	// admin.GET("/user/:id", hndler.HandleGetUserById)
	// admin.POST("/user", hndler.HandleInsertUser)
	// admin.PUT("/user", hndler.HandleUpdateUser)
	// admin.DELETE("/user/:id", hndler.HandleDeleteUser)

	return router
}
