package controllers

import (
	"net/http"

	"dota-esport/services"

	"github.com/gin-gonic/gin"
)

func SyncLeagues(c *gin.Context) {
	err := services.LeagueService.FetchAndSaveLeagues()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tier 1 leagues synced!"})
}

func GetLeagues(c *gin.Context) {
	leagues, err := services.LeagueService.GetLeagues()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch leagues"})
		return
	}

	c.JSON(http.StatusOK, leagues)
}
