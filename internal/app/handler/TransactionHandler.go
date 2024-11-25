package handler

import (
	"github.com/gin-gonic/gin"
	"mncTest/internal/app/dto"
	"mncTest/internal/app/services"
	"net/http"
)

type TransactionHandler struct {
}

func (a *TransactionHandler) CreateTransactions(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.JSON(http.StatusBadRequest, dto.CommonResponse{
			Status:  http.StatusBadRequest,
			Message: "Authorization header is missing",
			Data:    nil,
		})
		return
	}

	var request dto.TransactionRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.CommonResponse{
			Status:  http.StatusInternalServerError,
			Message: "Error unknown",
			Data:    nil,
		})
		return
	}
	transaction, err := services.CreateTransaction(authHeader, request.MerchantId, request.Amount)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.CommonResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusCreated, dto.CommonResponse{
		Status:  http.StatusCreated,
		Message: "Successfully created transaction",
		Data:    transaction,
	})
}

func (a *TransactionHandler) GetAllTransactions(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.JSON(http.StatusBadRequest, dto.CommonResponse{
			Status:  http.StatusBadRequest,
			Message: "Authorization header is missing",
			Data:    nil,
		})
		return
	}
	transactions, err := services.GetAllTransactions(authHeader)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.CommonResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.CommonResponse{
		Status:  http.StatusOK,
		Message: "Successfully retrieved all transactions",
		Data:    transactions,
	})
}
