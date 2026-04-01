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

func (s *PromotionService) calculateWinRate(u User) int {
	return (u.GamesWon * 100) / u.GamesTotal
}

func (s *PromotionService) formatSummary(u User, winRate int) (string, error) {
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

		winRate := s.calculateWinRate(u)

		if winRate < 50 {
			continue
		}

		summary, err := s.formatSummary(u, winRate)
		if err != nil {
			report.ErrorCount++
			errs = append(errs, fmt.Errorf("skipped user %d: %w", u.ID, err))
			continue
		}

		report.TopPlayers = append(report.TopPlayers, summary)
	}

	return report, errors.Join(errs...)
}

func (s *PromotionService) CalculateAverageWinRate() (int, error) {
	users, err := s.repo.GetAllUsers()
	if err != nil {
		return 0, fmt.Errorf("failed to fetch users from database: %w", err)
	}

	if len(users) == 0 {
		return 0, nil
	}

	totalWinRate := 0
	for _, u := range users {
		totalWinRate += s.calculateWinRate(u)
	}

	return totalWinRate / len(users), nil
}
