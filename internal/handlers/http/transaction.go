package http

import (
	"github.com/gin-gonic/gin"
	"github.com/rtzgod/EWallet/internal/handlers"
	"net/http"
)

type receiverWalletForm struct {
	ReceiverId string  `json:"to"  binding:"required"`
	Amount     float64 `json:"amount" binding:"required"`
}

// @Summary SendMoney
// @Tags Transaction
// @Description Sends money from one wallet to another
// @Accept json
// @Produce json
// @Param walletId path string true "WalletId"
// @Param transaction body receiverWalletForm true "Transaction"
// @Router /api/v1/wallet/{walletId}/send [post]
func (h *Handler) sendMoney(c *gin.Context) {
	var input receiverWalletForm
	senderId := c.Param("walletId")
	if err := c.BindJSON(&input); err != nil {
		handlers.NewErrorResponse(c, http.StatusBadRequest, "json body incorrect")
		return
	}

	if senderId == input.ReceiverId {
		handlers.NewErrorResponse(c, http.StatusBadRequest, "self transaction")
		return
	}

	senderWallet, err := h.services.Wallet.GetById(senderId)
	if err != nil {
		handlers.NewErrorResponse(c, http.StatusNotFound, "sender Wallet not found")
		return
	}
	if _, err := h.services.Wallet.GetById(input.ReceiverId); err != nil {
		handlers.NewErrorResponse(c, http.StatusBadRequest, "receiver Wallet not found")
		return
	}
	if senderWallet.Balance < input.Amount {
		handlers.NewErrorResponse(c, http.StatusBadRequest, "sender wallet balance not enough")
		return
	}
	if err := h.services.Update(senderId, input.ReceiverId, input.Amount); err != nil {
		handlers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if err := h.services.Transaction.Create(senderId, input.ReceiverId, input.Amount); err != nil {
		handlers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, handlers.StatusResponse{
		"ok",
	})
}

// @Summary GetHistory
// @Tags Transaction
// @Description returns history of transactions of wallet by id
// @Produce json
// @Param walletId path string true "WalletId"
// @Router /api/v1/wallet/{walletId}/history [get]
func (h *Handler) getHistory(c *gin.Context) {
	id := c.Param("walletId")
	if _, err := h.services.Wallet.GetById(id); err != nil {
		handlers.NewErrorResponse(c, http.StatusNotFound, "wallet not found")
		return
	}
	transactions, err := h.services.Transaction.GetAllById(id)
	if err != nil {
		handlers.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, transactions)
}
