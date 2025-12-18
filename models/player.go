package models

type Player struct {
	AccountID int    `json:"account_id"`
	TeamID    int    `json:"team_id"`
	Name      string `json:"name"`
	Country   string `json:"country"`
	IsLocked  bool   `json:"is_locked"`
	IsPro     bool   `json:"is_pro"`
	IsCurrent bool   `json:"is_current_team_member"`
}
