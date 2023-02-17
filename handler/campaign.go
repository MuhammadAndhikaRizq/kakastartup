package handler

import (
	"kakastartup/campaign"
	"kakastartup/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandller struct {
	service campaign.Service
}

// object dari campaignhandler
func NewCampaignHandler(service campaign.Service) *campaignHandller {
	return &campaignHandller{service}
}

// api/v1/campaigns
func (h *campaignHandller) GetCampaigns(c *gin.Context) {
	//Convert to int
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("Error to GetCampaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of Campaigns", http.StatusOK, "succes", campaigns)
	c.JSON(http.StatusOK, response)
}
