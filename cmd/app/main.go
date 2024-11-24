package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"mncTest/internal/app/handler"
	"os"
	"time"
)

func main() {
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
	router.POST("/login", authController.Login)
	router.POST("/logout", authController.Logout)
	router.POST("/transaction", transactionController.CreateTransactions)
	router.GET("/transaction", transactionController.GetAllTransactions)
	router.Run(":8080")
}
