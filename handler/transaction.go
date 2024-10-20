package handler

import (
	"net/http"
	"tiny-donate/helper"
	"tiny-donate/transaction"
	"tiny-donate/user"

	"github.com/gin-gonic/gin"
)

// parameter di uri
// tangkap parameter mapping input struct (pake should bind uri)
// panggil service, input struct sebagai parameternya
// service punya campaign_id

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func(h *transactionHandler) GetCampaignTransactions(c *gin.Context) {
	// tangkap inputnya dulu
	var input transaction.GetCampaignTransactionInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	transactions, err := h.service.GetTransactionsByCampaignID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Get campaign's transactions success", http.StatusOK, "success", transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, response)

}
