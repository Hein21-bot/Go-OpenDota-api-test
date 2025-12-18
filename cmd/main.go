package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"dota-esport/config"
	"dota-esport/cron"
	"dota-esport/db"
	"dota-esport/routes"
	"dota-esport/services"
)

func main() {
	config.LoadConfig()
	db.ConnectDB()

	r := gin.Default()
	routes.RegisterRoutes(r)

	cron.StartMatchCron(func() {
		services.MatchService.FetchMatchesByLeague(18629)
	})

	cron.UpdateUpcomingToLiveCron(func() {
		services.MatchService.UpdateUpcomingToLive()
		services.MatchService.UpdateLiveToFinished()
	})

	log.Println("Server running on :8080")
	r.Run(":8080")
}
