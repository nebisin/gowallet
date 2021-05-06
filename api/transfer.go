package api

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/nebisin/gowallet/db/model"
	"github.com/nebisin/gowallet/token"
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
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}

	fromAccount, valid := server.validAccount(c, arg.FromAccountID, req.Currency)
	if !valid {
		return
	}

	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	if fromAccount.Owner != authPayload.Username {
		err := errors.New("from account doesn't belong to the authenticated user")
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	_, valid = server.validAccount(c, arg.ToAccountID, req.Currency)
	if !valid {
		return
	}

	result, err := server.store.TransferTx(arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (server *Server) validAccount(c *gin.Context, accountID uint64, currency string) (model.Account, bool) {
	account, err := server.store.GetAccount(accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return account, false
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return account, false
	}

	if account.Currency != currency {
		c.JSON(http.StatusBadRequest, "Account currency is not match")
		return account, false
	}

	return account, true
}
