package services

import (
	"errors"

	"dota-esport/clients"
	"dota-esport/models"
	"dota-esport/repositories"
)

type leagueService struct {
	client *clients.OpenDotaClient
}

var LeagueService = &leagueService{
	client: clients.NewOpenDotaClient(),
}

func (s *leagueService) FetchAndSaveLeagues() error {
	var leagues []models.League
	if err := s.client.Get("/leagues", &leagues); err != nil {
		return err
	}
	if len(leagues) == 0 {
		return errors.New("no leagues returned")
	}

	for _, l := range leagues {
		if l.Tier == "premium" { // tier 1
			_ = repositories.LeagueRepository.SaveLeague(l)
		}
	}
	return nil
}

func (s *leagueService) GetLeagues() ([]models.League, error) {
	return repositories.LeagueRepository.GetLeagues()
}
