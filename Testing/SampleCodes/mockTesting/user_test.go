package user

import "errors"

type MockUserRepository struct {
	// You can customize the behavior here, e.g. with a map or a function field
	Users map[int]*User
	Err   error
}

func (m *MockUserRepository) FindByID(id int) (*User, error) {
	if m.Err != nil {
		return nil, m.Err
	}

	user, ok := m.Users[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func TestGetUserName_Success(t *testing.T) {
	mockRepo := &MockUserRepository{
		Users: map[int]*User{
			1: {ID: 1, Name: "Alice"},
		},
	}

	service := NewUserService(mockRepo)
	name, err := service.GetUserName(1)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if name != "Alice" {
		t.Errorf("expected 'Alice', got %s", name)
	}
}

func TestGetUserName_UserNotFound(t *testing.T) {
	mockRepo := &MockUserRepository{
		Users: map[int]*User{}, // empty map means no users found
	}

	service := NewUserService(mockRepo)
	_, err := service.GetUserName(1)

	if err == nil {
		t.Fatal("expected error for missing user, got nil")
	}
}

func TestGetUserName_ErrorFromRepo(t *testing.T) {
	mockRepo := &MockUserRepository{
		Err: errors.New("database error"),
	}

	service := NewUserService(mockRepo)
	_, err := service.GetUserName(1)

	if err == nil {
		t.Fatal("expected error from repo, got nil")
	}
}
