package routes

import (
	"dota-esport/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterLeagueRoutes(r *gin.Engine) {
	league := r.Group("/api/league")
	{
		league.GET("/sync", controllers.SyncLeagues)
		league.GET("/", controllers.GetLeagues)
	}
}
