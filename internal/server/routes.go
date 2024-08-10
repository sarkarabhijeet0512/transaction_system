package server

import (
	"transaction_system/internal/server/mw"

	"github.com/gin-gonic/gin"
)

func v1Routes(router *gin.RouterGroup, o *Options) {
	r := router.Group("/v1/")
	// transaction apis
	r.Use(mw.ErrorHandlerX(o.Log))
	router.PUT("/transactionservice/transaction/:transaction_id", o.TransactionHandler.CreateTransaction)
	router.GET("/transactionservice/transaction/:transaction_id", o.TransactionHandler.GetTransactionByID)
	router.GET("/transactionservice/types/:type", o.TransactionHandler.GetTransactionsByType)
	router.GET("/transactionservice/sum/:transaction_id", o.TransactionHandler.GetSumByTransactionID)
}
