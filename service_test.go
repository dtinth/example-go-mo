package main

import (
	"errors"
	"strings"
	"testing"
)

func TestPromotionService_DatabaseFailure(t *testing.T) {
	// Arrange
	repo := &MockUserRepository{
		MockErr: errors.New("connection timeout"),
	}
	service := NewPromotionService(repo)

	// Act
	_, err := service.GenerateReport()

	// Assert
	if err == nil || !strings.Contains(err.Error(), "failed to fetch users") {
		t.Fatalf("Expected database failure error, got: %v", err)
	}
}

func TestPromotionService_GenerateReport(t *testing.T) {
	// Arrange
	repo := &MockUserRepository{
		MockData: []User{
			{ID: 1, Username: "alice", GamesWon: 8, GamesTotal: 10},
			{ID: 2, Username: "bob", GamesWon: 2, GamesTotal: 10},
			{ID: 3, Username: "", GamesWon: 9, GamesTotal: 10},
		},
	}
	service := NewPromotionService(repo)

	// Act
	report, err := service.GenerateReport()

	// Assert
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

func TestPromotionService_CalculateAverageWinRate_Success(t *testing.T) {
	// Arrange
	repo := &MockUserRepository{
		MockData: []User{
			{ID: 1, Username: "alice", GamesWon: 8, GamesTotal: 10}, // Win rate: 80%
			{ID: 2, Username: "bob", GamesWon: 2, GamesTotal: 10},   // Win rate: 20%
		},
	}
	service := NewPromotionService(repo)

	// Act
	avg, err := service.CalculateAverageWinRate()

	// Assert
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	expectedAvg := 50
	if avg != expectedAvg {
		t.Errorf("Expected average win rate of %d%%, got %d%%", expectedAvg, avg)
	}
}

func TestPromotionService_CalculateAverageWinRate_DatabaseFailure(t *testing.T) {
	// Arrange
	repo := &MockUserRepository{
		MockErr: errors.New("database connection lost"),
	}
	service := NewPromotionService(repo)

	// Act
	_, err := service.CalculateAverageWinRate()

	// Assert
	if err == nil {
		t.Fatal("Expected an error due to database failure, got nil")
	}
	expectedErrSnippet := "failed to fetch users from database"
	if !strings.Contains(err.Error(), expectedErrSnippet) {
		t.Errorf("Error context missing. \nExpected it to contain: %q \nGot: %q", expectedErrSnippet, err.Error())
	}
}
