package routes

import (
	"dota-esport/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterPlayerRoutes(r *gin.Engine) {
	group := r.Group("/api/players")
	{
		group.GET("/sync/players/:teamId", controllers.Players.SyncPlayers)
		group.GET("/teams/:teamId/players", controllers.Players.GetPlayersByTeam)
	}
}
