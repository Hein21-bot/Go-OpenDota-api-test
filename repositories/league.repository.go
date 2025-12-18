package repositories

import (
	"dota-esport/db"
	"dota-esport/models"
	"log"
)

type leagueRepo struct{}

var LeagueRepository = &leagueRepo{}

func (r *leagueRepo) SaveLeague(l models.League) error {
	result := db.DB.Exec(`
        INSERT INTO leagues (id, name, tier, region)
        VALUES ($1, $2, $3, $4)
        ON CONFLICT (id) DO UPDATE SET
            name = EXCLUDED.name,
            tier = EXCLUDED.tier,
            region = EXCLUDED.region;
    `, l.ID, l.Name, l.Tier, l.Region)

	if result.Error != nil {
		log.Println("Insert error:", result.Error)
	} else if result.RowsAffected == 0 {
		log.Println("No rows affected")
	} else {
		log.Println("Inserted league:", l.ID)
	}

	return result.Error
}

func (r *leagueRepo) GetLeagues() ([]models.League, error) {
	var leagues []models.League
	result := db.DB.Find(&leagues)
	return leagues, result.Error
}
