package services

import (
	"context"

	"github.com/dbzyuzin/swagger/internal/models"
)

type UsersRepository interface {
	AddUser(models.User)
	FindUser(string) bool
}

type Service struct {
	repo UsersRepository
}

func New(repo UsersRepository) Service {
	return Service{
		repo: repo,
	}
}

func (s *Service) CreateNewUser(ctx context.Context, user models.User) error {
	s.repo.AddUser(user)

	return nil
}

func (s *Service) UserExists(ctx context.Context, name string) (bool, error) {
	ok := s.repo.FindUser(name)
	return ok, nil
}
