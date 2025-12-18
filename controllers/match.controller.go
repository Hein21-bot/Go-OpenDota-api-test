package controllers

import (
	"net/http"
	"strconv"

	"dota-esport/services"

	"github.com/gin-gonic/gin"
)

func FetchMatchesByLeague(c *gin.Context) {
	leagueID, _ := strconv.Atoi(c.Param("leagueId"))

	err := services.MatchService.FetchMatchesByLeague(leagueID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch matches"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Matches synced"})
}

func GetMatches(c *gin.Context) {
	matches, err := services.MatchService.GetMatches()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load matches"})
		return
	}

	c.JSON(http.StatusOK, matches)
}

func GetMatchesByLeague(c *gin.Context) {
	leagueID, _ := strconv.Atoi(c.Param("leagueId"))
	matches, err := services.MatchService.GetMatchesByLeague(leagueID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load matches"})
		return
	}

	c.JSON(http.StatusOK, matches)
}

func GetMatchesByTeam(c *gin.Context) {
	teamID, _ := strconv.Atoi(c.Param("teamId"))
	matches, err := services.MatchService.GetMatchesByTeam(teamID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed"})
		return
	}

	c.JSON(http.StatusOK, matches)
}

func GetTodayMatches(c *gin.Context) {
	matches, err := services.MatchService.GetTodayMatches()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, matches)
}

func GetUpcomingMatches(c *gin.Context) {
	filter := c.DefaultQuery("filter", "week")
	matches, err := services.MatchService.GetUpcoming(filter)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, matches)
}

func GetLiveMatches(c *gin.Context) {
	matches, _ := services.MatchService.GetLive()
	c.JSON(http.StatusOK, matches)
}

func GetFinishedMatches(c *gin.Context) {
	matches, _ := services.MatchService.GetFinished()
	c.JSON(http.StatusOK, matches)
}
