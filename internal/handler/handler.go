package handler

import (
	"context"
	"net/http"
	"notegin/internal/database"
	"notegin/internal/models"

	"github.com/gin-gonic/gin"
)

type handler struct {
	userstorer database.UserStorer
}

func NewHandler(userstorer database.UserStorer) handler {
	return handler{
		userstorer: userstorer,
	}
}

func (h handler) HandleGetNotes(ctx *gin.Context) {
	var pCtx, cancel = context.WithCancel(context.Background())
	defer cancel()

	h.userstorer.InsertUser(pCtx, models.User{})
	ctx.JSON(http.StatusOK, gin.H{
		"result": "there is no record",
	})
	return
}
