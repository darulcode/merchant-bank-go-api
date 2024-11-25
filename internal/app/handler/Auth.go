package handler

import (
	"github.com/gin-gonic/gin"
	"mncTest/internal/app/dto"
	"mncTest/internal/app/pkg/token"
	"mncTest/internal/app/services"
	"net/http"
	"time"
)

type AuthHandler struct {
}

func (a *AuthHandler) Register(ctx *gin.Context) {
	var req = dto.RegisterRequest{}

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	_, err = services.RegisterService(req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusOK, dto.CommonResponse{
			Status:  http.StatusOK,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusCreated, dto.CommonResponse{
		Status:  http.StatusCreated,
		Message: "Register Successfully",
		Data:    nil,
	})
}

func (a *AuthHandler) Login(ctx *gin.Context) {
	var req = dto.LoginRequest{}

	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	login, err := services.LoginService(req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	generateToken, err := token.GenerateToken(&token.PayloadToken{AuthId: login.Id, Exp: time.Now().Add(10 * time.Minute)})
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, dto.CommonResponse{
		Status:  http.StatusOK,
		Message: "LoginService Successfully",
		Data: dto.LoginResponse{
			Token: generateToken,
		},
	})
}

func (a *AuthHandler) Logout(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	_, err := services.LogoutService(authHeader)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Unauthorized",
			"data":    "",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Logout successful",
		"data":    "",
	})
}
