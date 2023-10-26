package handler

import (
	"net/http"
	"notegin/internal/database"

	"github.com/gin-gonic/gin"
)

type handler struct {
	database.Database
}

func NewHandler(db database.Database) handler {
	return handler{
		Database: db,
	}
}

func (h handler) HandleGetNotes(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"result": "there is no record",
	})
	return
}
