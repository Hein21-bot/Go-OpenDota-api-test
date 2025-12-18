package repositories

import (
	"dota-esport/db"
	"dota-esport/models"
	"dota-esport/utils"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type matchRepo struct{}

var MatchRepository = &matchRepo{}

func (r *matchRepo) SaveMatch(m models.Match) error {
	status := utils.CalculateMatchStatus(m.StartTime, m.Duration)
	result := db.DB.Exec(`
		INSERT INTO matches (id, league_id, radiant_id, dire_id, radiant_name, dire_name, start_time, duration, radiant_score, dire_score, status)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8, $9, $10, $11)
		ON CONFLICT (id) DO UPDATE SET
			league_id = EXCLUDED.league_id,
			radiant_id = EXCLUDED.radiant_id,
			dire_id = EXCLUDED.dire_id,
			radiant_name = EXCLUDED.radiant_name,
			dire_name = EXCLUDED.dire_name,
			start_time = EXCLUDED.start_time,
			duration = EXCLUDED.duration,
			radiant_score = EXCLUDED.radiant_score,
			dire_score = EXCLUDED.dire_score,
			status = EXCLUDED.status
	`, m.ID, m.LeagueID, m.RadiantID, m.DireID, m.RadiantName, m.DireName, m.StartTime, m.Duration, m.RadiantScore, m.DireScore, status)

	if result.Error != nil {
		log.Println("Insert error:", result.Error)
	} else if result.RowsAffected == 0 {
		log.Println("No rows affected")
	} else {
		log.Println("Inserted matches:", m.ID)
	}

	return result.Error
}

func (r *matchRepo) GetMatches() ([]models.Match, error) {
	var matches []models.Match
	result := db.DB.Find(&matches)
	return matches, result.Error
}

func (r *matchRepo) GetMatchesByLeague(leagueID int) ([]models.Match, error) {
	var matches []models.Match
	result := db.DB.
		Where("league_id = ?", leagueID).
		Order("start_time DESC").
		Find(&matches)
	if result.Error != nil {
		return nil, result.Error
	}
	return matches, nil
}

func (r *matchRepo) GetMatchesByTeam(teamID int) ([]models.Match, error) {
	var matches []models.Match
	result := db.DB.
		Where("radiant_id = ? OR dire_id = ?", teamID, teamID).
		Order("start_time").
		Find(&matches)

	if result.Error != nil {
		return nil, result.Error
	}

	return matches, nil
}

func (r *matchRepo) GetMatchesBetween(startUnix, endUnix int64) ([]models.MatchResponse, error) {

	var rawMatches []models.Match
	result := db.DB.
		Where("start_time >= ? AND start_time < ?", startUnix, endUnix).
		Order("start_time ASC").
		Find(&rawMatches)

	if result.Error != nil {
		return nil, result.Error
	}

	var response []models.MatchResponse

	for _, m := range rawMatches {

		var league models.League
		db.DB.First(&league, m.LeagueID)

		// var radiant models.Team
		// db.DB.First(&radiant, m.RadiantID)

		// var dire models.Team
		// db.DB.First(&dire, m.DireID)

		radiant := getTeamOrPlaceholder(m.RadiantID)
		dire := getTeamOrPlaceholder(m.DireID)

		formatted := time.Unix(m.StartTime, 0).Format("2006-01-02 15:04")

		response = append(response, models.MatchResponse{
			ID:           int(m.ID),
			LeagueID:     int(m.LeagueID),
			LeagueName:   league.Name,
			RadiantID:    int(m.RadiantID),
			RadiantName:  radiant.Name,
			DireID:       m.DireID,
			DireName:     dire.Name,
			StartTime:    formatted,
			Duration:     m.Duration,
			RadiantScore: m.RadiantScore,
			DireScore:    m.DireScore,
			Status:       m.Status,
		})
	}

	return response, nil
}

func getTeamOrPlaceholder(id int) models.Team {
	var t models.Team
	if id == 0 {
		return models.Team{ID: 0, Name: "Unknown"}
	}

	err := db.DB.First(&t, id).Error
	if err != nil {
		return models.Team{ID: id, Name: "Unknown"}
	}
	return t
}

func (r *matchRepo) FindByStatus(status string) ([]models.Match, error) {
	var matches []models.Match
	err := db.DB.Where("status = ?", status).
		Order("start_time ASC").
		Find(&matches).Error
	return matches, err
}

func (r *matchRepo) FindUpcomingRange(from, to int64) ([]models.Match, error) {
	var matches []models.Match
	err := db.DB.
		Where("start_time Between ? AND ? AND status = ?", from, to, "upcoming").
		Order("start_time ASC").
		Find(&matches).Error
	return matches, err
}

func (r *matchRepo) UpdateUpcomingToLive() error {
	return db.DB.Exec(`
		UPDATE matches
		SET status = 'LIVE'
		WHERE status = 'UPCOMING'
		  AND start_time <= EXTRACT(EPOCH FROM NOW())
	`).Error
}

func (r *matchRepo) UpdateLiveToFinished() error {
	var matches []struct {
		ID int64
	}
	db.DB.
		Table("matches").
		Select("id").
		Where("status = ?", "LIVE").
		Scan(&matches)

	for _, m := range matches {
		url := fmt.Sprintf("https://api.opendota.com/api/matches/%d", m.ID)
		resp, err := http.Get(url)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		var data struct {
			Duration int `json:"duration"`
		}

		body, _ := io.ReadAll(resp.Body)
		json.Unmarshal(body, &data)

		if data.Duration > 0 {
			db.DB.Exec(`
				UPDATE matches
				SET duration = ?, status = 'FINISHED'
				WHERE id = ?
			`, data.Duration, m.ID)
		}
	}
	return nil
}
