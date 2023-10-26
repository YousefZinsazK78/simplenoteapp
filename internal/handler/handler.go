package handler

import (
	"context"
	"net/http"
	"notegin/internal/database"
	"notegin/internal/models"
	"strconv"

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

func (h handler) HandleGetUsers(ctx *gin.Context) {
	var pCtx, cancel = context.WithCancel(context.Background())
	defer cancel()

	users, err := h.userstorer.ViewUsers(pCtx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": users,
	})
	return
}

func (h handler) HandleGetUserByUsername(ctx *gin.Context) {
	var pCtx, cancel = context.WithCancel(context.Background())
	defer cancel()

	user, err := h.userstorer.ViewUserByUsername(pCtx, ctx.Param("username"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": user,
	})
	return
}

func (h handler) HandleGetUserById(ctx *gin.Context) {
	var pCtx, cancel = context.WithCancel(context.Background())
	defer cancel()

	userid, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	user, err := h.userstorer.ViewUserByID(pCtx, userid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": user,
	})
	return
}

func (h handler) HandleInsertUser(ctx *gin.Context) {
	var pCtx, cancel = context.WithCancel(context.Background())
	defer cancel()

	var usermodel models.User
	if err := ctx.ShouldBind(&usermodel); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	err := h.userstorer.InsertUser(pCtx, usermodel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": "insert successfully ✅",
	})
	return
}

func (h handler) HandleUpdateUser(ctx *gin.Context) {
	var pCtx, cancel = context.WithCancel(context.Background())
	defer cancel()

	var updateUserModel models.UpdateUserParams
	if err := ctx.ShouldBind(&updateUserModel); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	err := h.userstorer.UpdateUser(pCtx, updateUserModel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": "update successfully ✅",
	})
	return
}

func (h handler) HandleDeleteUser(ctx *gin.Context) {
	var pCtx, cancel = context.WithCancel(context.Background())
	defer cancel()

	userid, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	err = h.userstorer.DeleteUser(pCtx, userid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": "delete successfully ✅",
	})
	return
}
