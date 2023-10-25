package handler

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	database
}

func NewHandler(db *sql.DB) handler {
	return handler{
		database: db,
	}
}

func (h handler) HandleGetNotes(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"result": "there is no record",
	})
	return
}
