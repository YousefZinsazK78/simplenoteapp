package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct{}

func NewHandler() handler {
	return handler{}
}

func (h handler) HandleGetNotes(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"result": "there is no record",
	})
	return
}
