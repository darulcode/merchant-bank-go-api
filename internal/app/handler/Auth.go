package handler

import (
	"github.com/gin-gonic/gin"
	"mncTest/internal/app/dto"
	"mncTest/internal/app/pkg/token"
	"mncTest/internal/app/services"
	"net/http"
)

type AuthHandler struct {
}

func (a *AuthHandler) Login(ctx *gin.Context) {
	var req = dto.LoginRequest{}

	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	login, err := services.Login(req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	generateToken, err := token.GenerateToken(&token.PayloadToken{AuthId: login.Id})
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, dto.CommonResponse{
		Status:  http.StatusOK,
		Message: "Login Successfully",
		Data: dto.LoginResponse{
			Token: generateToken,
		},
	})
}

func (a *AuthHandler) Logout(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	_, err := services.Logout(authHeader)
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
