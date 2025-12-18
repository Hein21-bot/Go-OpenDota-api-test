package routes

import (
	"dota-esport/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterMatchRoutes(r *gin.Engine) {
	group := r.Group("/api/matches")
	{
		group.GET("/", controllers.GetMatches)
		group.GET("/league/:leagueId", controllers.GetMatchesByLeague)
		group.GET("/team/:teamId", controllers.GetMatchesByTeam)

		group.GET("/sync/league/:leagueId", controllers.FetchMatchesByLeague)

		group.GET("/today", controllers.GetTodayMatches)

		group.GET("/upcoming", controllers.GetUpcomingMatches)
		group.GET("/live", controllers.GetLiveMatches)
		group.GET("/finished", controllers.GetFinishedMatches)
	}
}
