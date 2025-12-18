package repositories

import (
	"dota-esport/db"
	"dota-esport/models"
	"log"
)

type teamRepo struct{}

var TeamRepository = &teamRepo{}

// Save or Update Team (UPSERT)
func (r *teamRepo) SaveTeam(t models.Team) error {
	result := db.DB.Exec(`
		INSERT INTO teams (id, name, tag, logo_url, rating, region)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (id) DO UPDATE SET
			name = EXCLUDED.name,
			tag = EXCLUDED.tag,
			logo_url = EXCLUDED.logo_url,
			rating = EXCLUDED.rating,
			region = EXCLUDED.region
	`, t.ID, t.Name, t.Tag, t.LogoURL, t.Rating, t.Region)

	if result.Error != nil {
		log.Println("Insert error:", result.Error)
	} else if result.RowsAffected == 0 {
		log.Println("No rows affected")
	} else {
		log.Println("Inserted Team:", t.ID)
	}

	return result.Error
}

func (r *teamRepo) GetTeams() ([]models.Team, error) {
	var teams []models.Team
	result := db.DB.Find(&teams)
	return teams, result.Error
}

func (r *teamRepo) GetTeam(id int) (*models.Team, error) {
	var t models.Team
	result := db.DB.First(&t, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &t, nil
}
