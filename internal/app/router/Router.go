package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"mncTest/internal/app/handler"
	"os"
	"time"
)

func SetupRouter() *gin.Engine {
	logFile, err := os.OpenFile("gin.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Could not open log file")
	}
	router := gin.New()
	router.Use(gin.LoggerWithWriter(logFile))
	router.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		status := c.Writer.Status()
		log.Printf("[Custom Log] Method: %s, Path: %s, Status: %d, Duration: %v\n", c.Request.Method, c.Request.URL.Path, status, duration)
	})

	authController := &handler.AuthHandler{}
	transactionController := &handler.TransactionHandler{}
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/logout", authController.Logout)
	}

	transactionRoutes := router.Group("/transaction")
	{
		transactionRoutes.POST("", transactionController.CreateTransactions)
		transactionRoutes.GET("", transactionController.GetAllTransactions)
	}

	return router
}
