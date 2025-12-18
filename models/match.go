package models

type Match struct {
	ID           int64  `json:"match_id"`
	LeagueID     int    `json:"leagueid"`
	RadiantID    int    `json:"radiant_team_id"`
	DireID       int    `json:"dire_team_id"`
	RadiantName  string `json:"radiant_team_name"`
	DireName     string `json:"dire_team_name"`
	StartTime    int64  `json:"start_time"`
	Duration     int    `json:"duration"`
	RadiantScore int    `json:"radiant_score"`
	DireScore    int    `json:"dire_score"`
	Status       string `json:"status"`
	UpdatedAt    string `json:"updated_at"`
}
