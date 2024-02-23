package services

import (
	"context"
	"fmt"

	"github.com/dbzyuzin/swagger/internal/models"
	"github.com/dbzyuzin/swagger/internal/pkg/tracing"
)

type UsersRepository interface {
	AddUser(models.User)
	FindUser(context.Context, string) bool
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
	fmt.Println("from service", tracing.Get(ctx))
	ok := s.repo.FindUser(ctx, name)
	return ok, nil
}

func (s *Service) NotifyTodayWather(context.Context) error {
	fmt.Println("NOTIFY")
	return nil
}
