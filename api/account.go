package api

import (
	"github.com/gin-gonic/gin"
	"github.com/nebisin/gowallet/db/model"
	"net/http"
)

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR TRY"`
}

func (s *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := model.CreateAccountPayload{
		Owner: req.Owner,
		Balance: 0,
		Currency: req.Currency,
	}

	account, err := s.store.CreateAccount(arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, account)
}
