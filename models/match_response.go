package models

type MatchResponse struct {
	ID           int    `json:"id"`
	LeagueID     int    `json:"league_id"`
	LeagueName   string `json:"league_name"`
	RadiantID    int    `json:"radiant_id"`
	RadiantName  string `json:"radiant_name"`
	DireID       int    `json:"dire_id"`
	DireName     string `json:"dire_name"`
	StartTime    string `json:"start_time"` // formatted
	Duration     int    `json:"duration"`
	Status       string `json:"status"`
	RadiantScore int    `json:"radiant_score"`
	DireScore    int    `json:"dire_score"`
}
