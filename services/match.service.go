package services

import (
	"fmt"

	"dota-esport/clients"
	"dota-esport/models"
	"dota-esport/repositories"
	"dota-esport/utils"
	"errors"
	"time"
)

type matchService struct {
	client *clients.OpenDotaClient
}

var MatchService = &matchService{
	client: clients.NewOpenDotaClient(),
}

func (s *matchService) FetchMatchesByLeague(leagueID int) error {
	var matches []models.Match
	path := fmt.Sprintf("/leagues/%d/matches", leagueID)

	if err := s.client.Get(path, &matches); err != nil {
		return err
	}
	if len(matches) == 0 {
		return errors.New("no matches returned")
	}
	for _, m := range matches {
		repositories.MatchRepository.SaveMatch(m)
	}

	return nil
}

func (s *matchService) GetMatches() ([]models.Match, error) {
	return repositories.MatchRepository.GetMatches()
}

func (s *matchService) GetMatchesByLeague(id int) ([]models.Match, error) {
	return repositories.MatchRepository.GetMatchesByLeague(id)
}

func (s *matchService) GetMatchesByTeam(id int) ([]models.Match, error) {
	return repositories.MatchRepository.GetMatchesByTeam(id)
}

func (s *matchService) GetTodayMatches() (interface{}, error) {

	now := time.Now().UTC()

	start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	end := start.Add(24 * time.Hour)

	return repositories.MatchRepository.GetMatchesBetween(start.Unix(), end.Unix())
}

func (s *matchService) GetUpcoming(filter string) ([]models.Match, error) {
	var from, to int64
	switch filter {
	case "today":
		from, to = utils.TodayRange()
	case "tomorrow":
		from, to = utils.TomorrowRange()
	default:
		from, to = utils.WeekRange()
	}
	return repositories.MatchRepository.FindUpcomingRange(from, to)
}

func (s *matchService) GetLive() ([]models.Match, error) {
	return repositories.MatchRepository.FindByStatus("live")
}

func (s *matchService) GetFinished() ([]models.Match, error) {
	return repositories.MatchRepository.FindByStatus("finished")
}

func (s *matchService) UpdateUpcomingToLive() error {
	return repositories.MatchRepository.UpdateUpcomingToLive()
}

func (s *matchService) UpdateLiveToFinished() error {
	return repositories.MatchRepository.UpdateLiveToFinished()
}
