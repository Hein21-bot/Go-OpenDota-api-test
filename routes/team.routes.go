package routes

import (
	"dota-esport/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterTeamRoutes(r *gin.Engine) {
	group := r.Group("/api/teams")
	{
		group.GET("/", controllers.GetTeams)
		group.GET("/:id", controllers.GetTeam)
		group.GET("/sync", controllers.FetchTeams)
	}
}
