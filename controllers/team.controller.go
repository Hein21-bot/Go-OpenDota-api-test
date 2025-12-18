package controllers

import (
	"net/http"
	"strconv"

	"dota-esport/services"

	"github.com/gin-gonic/gin"
)

func FetchTeams(c *gin.Context) {
	if err := services.TeamService.FetchTeams(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch teams"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Teams synced successfully"})
}

func GetTeams(c *gin.Context) {
	teams, err := services.TeamService.GetTeams()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load teams"})
		return
	}

	c.JSON(http.StatusOK, teams)
}

func GetTeam(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	team, err := services.TeamService.GetTeam(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Team not found"})
		return
	}

	c.JSON(http.StatusOK, team)
}
