package services

import (
	"dota-esport/clients"
	"dota-esport/models"
	"dota-esport/repositories"
	"fmt"
)

type teamService struct {
	client *clients.OpenDotaClient
}

var TeamService = &teamService{
	client: clients.NewOpenDotaClient(),
}

func (s *teamService) FetchTeams() error {
	const PAGE_LIMIT = 100

	for page := range PAGE_LIMIT {
		var teams []models.Team
		path := fmt.Sprintf("/teams?page=%d", page)

		if err := s.client.Get(path, &teams); err != nil {
			return err
		}
		// STOP when no more data
		if len(teams) == 0 {
			fmt.Println("No more teams. Stopping at page:", page)
			break
		}

		// SAVE ONLY REAL TEAMS
		for _, t := range teams {
			// if isRealTeam(t) {
			repositories.TeamRepository.SaveTeam(t)
			// }
		}
	}

	return nil
}

func (s *teamService) GetTeams() ([]models.Team, error) {
	return repositories.TeamRepository.GetTeams()
}

func (s *teamService) GetTeam(id int) (*models.Team, error) {
	return repositories.TeamRepository.GetTeam(id)
}
