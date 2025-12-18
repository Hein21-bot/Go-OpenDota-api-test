package services

import (
	"dota-esport/clients"
	"fmt"

	"dota-esport/models"
	"dota-esport/repositories"
)

type playerService struct {
	client *clients.OpenDotaClient
}

var PlayerService = &playerService{
	client: clients.NewOpenDotaClient(),
}

func (s *playerService) SyncPlayers(teamID int) error {
	var players []models.Player
	path := fmt.Sprintf("/teams/%d/players", teamID)

	if err := s.client.Get(path, &players); err != nil {
		return err
	}
	for _, p := range players {
		player := models.Player{
			AccountID: p.AccountID,
			TeamID:    teamID,
			Name:      p.Name,
			Country:   "not specify",
			IsPro:     true,
			IsLocked:  p.IsLocked,
			IsCurrent: p.IsCurrent,
		}
		repositories.PlayerRepository.SavePlayer(player)
	}
	return nil
}
