package cron

import (
	"log"

	"github.com/robfig/cron/v3"
)

func StartMatchCron(syncFunc func()) {
	c := cron.New()
	c.AddFunc("* * * * *", func() {
		log.Println("Running Nightly match sync")
		syncFunc()
	})
	c.Start()
}

func UpdateUpcomingToLiveCron(syncFunc func()) {
	c := cron.New()
	c.AddFunc("* * * * *", func() {
		log.Println(" upcoming → live (every minute)")
		log.Println("upcoming → live + live → finished + duration update")
		syncFunc()
	})
	c.Start()
}
