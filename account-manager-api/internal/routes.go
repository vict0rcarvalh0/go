package routes

import (
	"account-manager-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.POST("/accounts", handlers.CreateAccount)
	r.DELETE("/accounts/:id", handlers.DeleteAccount)
	r.POST("/accounts/:id/deposit", handlers.Deposit)
	r.POST("/accounts/:id/withdraw", handlers.Withdraw)
	r.POST("/transfer", handlers.Transfer)
	r.GET("/accounts", handlers.GetAccounts)

	return r
}
