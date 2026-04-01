package main

type MockUserRepository struct {
	MockData []User
	MockErr  error
}

func (m *MockUserRepository) GetAllUsers() ([]User, error) {
	if m.MockErr != nil {
		return nil, m.MockErr
	}
	return m.MockData, nil
}
