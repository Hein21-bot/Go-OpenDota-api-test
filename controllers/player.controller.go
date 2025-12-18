package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"dota-esport/repositories"
	"dota-esport/services"
)

type PlayerController struct{}

var Players = &PlayerController{}

func (c *PlayerController) SyncPlayers(ctx *gin.Context) {
	teamID, _ := strconv.Atoi(ctx.Param("teamId"))

	err := services.PlayerService.SyncPlayers(teamID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Players synced"})
}

func (c *PlayerController) GetPlayersByTeam(ctx *gin.Context) {
	teamID, _ := strconv.Atoi(ctx.Param("teamId"))

	players, err := repositories.PlayerRepository.GetPlayersByTeam(teamID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, players)
}
