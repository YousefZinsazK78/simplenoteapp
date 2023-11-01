package handler

import (
	"context"
	"net/http"
	"notegin/internal/models"
	"notegin/internal/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func (h handler) HandleTmplSignIn(ctx *gin.Context) {
	var pCtx, cancel = context.WithCancel(context.Background())
	defer cancel()

	var userParams models.UserParamsForm
	if err := ctx.ShouldBind(&userParams); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	//call db for these specific user
	var user, err = h.userstorer.ViewUserByUsername(pCtx, userParams.Username)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	//check for valid password
	var isValidPassword = utils.CompareAndVerifyPassword(user.Password, []byte(userParams.Password))
	if isValidPassword && user.Username == userParams.Username {
		generatedJWT := utils.GenerateJwt(user.ID)
		ctx.SetCookie("Authorization", generatedJWT, int(time.Hour*3), "localhost", "localhost", false, false)
		ctx.Redirect(http.StatusOK, "http://localhost:8000/")
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid credentials!"})
		return
	}
}

func (h handler) HandleTmplGetNotes(ctx *gin.Context) {
	var pCtx, cancel = context.WithCancel(context.Background())
	defer cancel()

	note, err := h.noteStorer.GetAll(pCtx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.HTML(http.StatusOK, "all.tmpl", gin.H{
		"notes": note,
	})
}

func (h handler) HandleTmplCreateNote(ctx *gin.Context) {
	var pCtx, cancel = context.WithCancel(context.Background())
	defer cancel()

	var notemodel models.NoteForm
	if err := ctx.ShouldBind(&notemodel); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	// userid, ok := ctx.Get("user_id")
	// if !ok {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "there's no userid"})
	// 	return
	// }
	// notemodel.UserID = int(userid.(float64))
	notemodel.UserID = 3
	err := h.noteStorer.Insert(pCtx, models.Note{
		Title:  notemodel.Title,
		Body:   notemodel.Body,
		UserID: notemodel.UserID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.Redirect(http.StatusOK, "http://localhost:8000/")
}
