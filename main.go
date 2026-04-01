package main

import "fmt"

func main() {
	repo := &MockUserRepository{
		MockData: []User{
			{ID: 1, Username: "alice", GamesWon: 8, GamesTotal: 10},
			{ID: 2, Username: "bob", GamesWon: 2, GamesTotal: 10},
			{ID: 3, Username: "", GamesWon: 9, GamesTotal: 10},
		},
	}

	service := NewPromotionService(repo)
	report, err := service.GenerateReport()

	fmt.Printf("Processed %d users. Found %d errors.\n", report.TotalProcessed, report.ErrorCount)
	if err != nil {
		fmt.Println("Job finished with accumulated errors:\n", err)
	}

	fmt.Println("\nTop Players:")
	for _, p := range report.TopPlayers {
		fmt.Println("-", p)
	}
}
