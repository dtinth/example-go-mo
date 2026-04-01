package main

type User struct {
	ID         int
	Username   string
	GamesWon   int
	GamesTotal int
}

type TopPlayerReport struct {
	TotalProcessed int
	TopPlayers     []string
	ErrorCount     int
}

type UserRepository interface {
	GetAllUsers() ([]User, error)
}