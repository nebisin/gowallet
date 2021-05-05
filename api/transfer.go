package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/nebisin/gowallet/db/model"
	"net/http"
)

type transferRequest struct {
	FromAccountID uint64 `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   uint64 `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,oneof=USD EUR TRY"`
}

func (server *Server) transfer(c *gin.Context) {
	var req transferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := model.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID: req.ToAccountID,
		Amount:  req.Amount,
	}

	if !server.validAccount(c, arg.FromAccountID, req.Currency) {
		return
	}
	if !server.validAccount(c, arg.ToAccountID, req.Currency) {
		return
	}

	result, err := server.store.TransferTx(arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (server *Server) validAccount(c *gin.Context, accountID uint64, currency string) bool {
	account, err := server.store.GetAccount(accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return false
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}

	if account.Currency != currency {
		c.JSON(http.StatusBadRequest, "Account currency is not match")
		return false
	}

	return true
}
