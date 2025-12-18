package models

type Team struct {
	ID            int     `gorm:"primaryKey" json:"team_id"`
	Name          string  `json:"name"`
	Tag           string  `json:"tag"`
	LogoURL       string  `json:"logo_url"`
	Rating        float64 `json:"rating"`
	Region        string  `json:"region"`
	Wins          int     `json:"wins"`
	Losses        int     `json:"losses"`
	LastMatchTime int     `json:"last_match_time"`
}
