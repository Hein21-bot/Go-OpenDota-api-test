package repositories

import (
	"dota-esport/db"
	"dota-esport/models"
	"log"
)

type playerRepo struct{}

var PlayerRepository = &playerRepo{}

func (r *playerRepo) SavePlayer(p models.Player) error {
	result := db.DB.Exec(`
		INSERT INTO players (account_id, team_id, name, country, is_locked, is_pro, is_current)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (account_id) DO UPDATE SET
			team_id = EXCLUDED.team_id,
			name = EXCLUDED.name,
			country = EXCLUDED.country,
			is_locked = EXCLUDED.is_locked,
			is_pro = EXCLUDED.is_pro,
			is_current = EXCLUDED.is_current
	`,
		p.AccountID,
		p.TeamID,
		p.Name,
		p.Country,
		p.IsLocked,
		p.IsPro,
		p.IsCurrent,
	)

	if result.Error != nil {
		log.Println("Insert error:", result.Error)
	} else if result.RowsAffected == 0 {
		log.Println("No rows affected")
	} else {
		log.Println("Inserted player:", p.AccountID)
	}

	return result.Error
}

func (r *playerRepo) GetPlayersByTeam(teamID int) ([]models.Player, error) {
	var player []models.Player
	result := db.DB.Find(&player)
	return player, result.Error
}
