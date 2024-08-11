package handler

import (
	"context"
	"strconv"
	"transaction_system/er"
	model "transaction_system/utils/models"

	"net/http"
	"transaction_system/pkg/transaction"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type TransactionHandler struct {
	log                *logrus.Logger
	transactionService *transaction.Service
}

func newTransactionHandler(
	log *logrus.Logger,
	transactionService *transaction.Service,
) *TransactionHandler {
	return &TransactionHandler{
		log:                log,
		transactionService: transactionService,
	}
}

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	var (
		err error
		res = model.GenericRes{}
		req = &transaction.Transaction{}
		ctx = context.Background()
	)
	defer func() {
		if err != nil {
			c.Error(err)
			h.log.WithField("info", err).Warn(err.Error())
		}
	}()

	if err = c.ShouldBind(&req); err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusBadRequest)
		return
	}
	req.TransactionID, _ = strconv.ParseInt(c.Param("transaction_id"), 10, 64)
	if err = h.transactionService.CreateTransaction(ctx, req); err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusUnprocessableEntity)
		return
	}

	res.Message = "Success"
	res.Success = true
	res.Data = req
	c.JSON(http.StatusOK, res)
}

func (h *TransactionHandler) GetTransactionByID(c *gin.Context) {
	var (
		err error
		res = model.GenericRes{}
		req = &transaction.Transaction{}
		ctx = context.Background()
	)

	defer func() {
		if err != nil {
			c.Error(err)
			h.log.WithField("info", err).Warn(err.Error())
		}
	}()

	if err = c.ShouldBind(&req); err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusBadRequest)
		return
	}
	id, _ := strconv.ParseInt(c.Param("transaction_id"), 10, 64)
	tx, err := h.transactionService.GetTransactionByID(ctx, id)
	if err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusUnprocessableEntity)
		return
	}

	res.Message = "Success"
	res.Success = true
	res.Data = tx
	c.JSON(http.StatusOK, res)
}

func (h *TransactionHandler) GetTransactionsByType(c *gin.Context) {
	var (
		err error
		res = model.GenericRes{}
		req = &transaction.Transaction{}
		ctx = context.Background()
	)
	defer func() {
		if err != nil {
			c.Error(err)
			h.log.WithField("info", err).Warn(err.Error())
		}
	}()

	if err = c.ShouldBind(&req); err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusBadRequest)
		return
	}
	txType := c.Param("type")
	ids, err := h.transactionService.GetTransactionsByType(ctx, txType)
	if err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusUnprocessableEntity)
		return
	}
	res.Message = "Success"
	res.Success = true
	res.Data = ids
	c.JSON(http.StatusOK, res)
}

func (h *TransactionHandler) GetSumByTransactionID(c *gin.Context) {
	var (
		err error
		res = model.GenericRes{}
		req = &transaction.Transaction{}
		ctx = context.Background()
	)
	defer func() {
		if err != nil {
			c.Error(err)
			h.log.WithField("info", err).Warn(err.Error())
		}
	}()

	if err = c.ShouldBind(&req); err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusBadRequest)
		return
	}
	id, _ := strconv.ParseInt(c.Param("transaction_id"), 10, 64)
	sum, err := h.transactionService.GetSumByTransactionID(ctx, id)
	if err != nil {
		err = er.New(err, er.UncaughtException).SetStatus(http.StatusUnprocessableEntity)
		return
	}
	res.Message = "Success"
	res.Success = true
	res.Data = sum
	c.JSON(http.StatusOK, res)
}
