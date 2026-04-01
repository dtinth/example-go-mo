package main

import (
	"errors"
	"strings"
	"testing"
)

func TestPromotionService_DatabaseFailure(t *testing.T) {
	repo := &MockUserRepository{
		MockErr: errors.New("connection timeout"),
	}
	service := NewPromotionService(repo)

	_, err := service.GenerateReport()
	if err == nil || !strings.Contains(err.Error(), "failed to fetch users") {
		t.Fatalf("Expected database failure error, got: %v", err)
	}
}

func TestPromotionService_GenerateReport(t *testing.T) {
	repo := &MockUserRepository{
		MockData: []User{
			{ID: 1, Username: "alice", GamesWon: 8, GamesTotal: 10},
			{ID: 2, Username: "bob", GamesWon: 2, GamesTotal: 10},
			{ID: 3, Username: "", GamesWon: 9, GamesTotal: 10},
		},
	}
	service := NewPromotionService(repo)

	report, err := service.GenerateReport()

	if len(report.TopPlayers) != 1 {
		t.Errorf("Expected 1 top player, got %d", len(report.TopPlayers))
	}
	if report.ErrorCount != 1 {
		t.Errorf("Expected 1 error, got %d", report.ErrorCount)
	}
	if err == nil || !strings.Contains(err.Error(), "skipped user 3: player is missing username") {
		t.Errorf("Expected contextual error for user 3, got: %v", err)
	}
}
