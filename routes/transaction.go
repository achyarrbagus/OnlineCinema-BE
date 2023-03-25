package routes

import (
	"backEnd/handlers"
	"backEnd/pkg/middleware"
	"backEnd/pkg/mysql"
	"backEnd/repositories"

	"github.com/labstack/echo/v4"
)

func TransactionRoutes(e *echo.Group) {
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository)

	// e.GET("/transaction/:id", middleware.Auth(h.GetTransaction))
	e.POST("/createtransaction", middleware.Auth(h.CreateTransaction))
	e.POST("/notification", h.Notification)
	e.GET("/transactions", middleware.Auth(h.FindTransaction))
	e.GET("/transaction-user", middleware.Auth(h.GetUserTransaction))
	// r.HandleFunc("/transaction/{id}", h.GetTransaction).Methods("GET")

}
