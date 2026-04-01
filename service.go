package main

import (
	"errors"
	"fmt"
)

type PromotionService struct {
	repo UserRepository
}

func NewPromotionService(repo UserRepository) *PromotionService {
	return &PromotionService{repo: repo}
}

func (s *PromotionService) CalculateWinRate(u User) int {
	return (u.GamesWon * 100) / u.GamesTotal
}

func (s *PromotionService) FormatSummary(u User, winRate int) (string, error) {
	if u.Username == "" {
		return "", errors.New("player is missing username")
	}
	return fmt.Sprintf("Player %s (ID: %d) - Win Rate: %d%%", u.Username, u.ID, winRate), nil
}

func (s *PromotionService) GenerateReport() (TopPlayerReport, error) {
	var report TopPlayerReport
	var errs []error

	users, err := s.repo.GetAllUsers()
	if err != nil {
		return report, fmt.Errorf("failed to fetch users from database: %w", err)
	}

	for _, u := range users {
		report.TotalProcessed++

		winRate := s.CalculateWinRate(u)

		if winRate < 50 {
			continue
		}

		summary, err := s.FormatSummary(u, winRate)
		if err != nil {
			report.ErrorCount++
			errs = append(errs, fmt.Errorf("skipped user %d: %w", u.ID, err))
			continue
		}

		report.TopPlayers = append(report.TopPlayers, summary)
	}

	return report, errors.Join(errs...)
}
