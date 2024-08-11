package server

import (
	"transaction_system/internal/server/mw"

	"github.com/gin-gonic/gin"
)

func v1Routes(router *gin.RouterGroup, o *Options) {
	r := router.Group("/v1/")
	// transaction apis
	r.Use(mw.ErrorHandlerX(o.Log))
	r.PUT("/transactionservice/transaction/:transaction_id", o.TransactionHandler.CreateTransaction)
	r.GET("/transactionservice/transaction/:transaction_id", o.TransactionHandler.GetTransactionByID)
	r.GET("/transactionservice/types/:type", o.TransactionHandler.GetTransactionsByType)
	r.GET("/transactionservice/sum/:transaction_id", o.TransactionHandler.GetSumByTransactionID)
}
