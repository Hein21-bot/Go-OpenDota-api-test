package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	RegisterLeagueRoutes(r)
	RegisterTeamRoutes(r)
	RegisterMatchRoutes(r)
	RegisterPlayerRoutes(r)
}
