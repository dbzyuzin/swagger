package services

import (
	"context"

	"github.com/dbzyuzin/swagger/internal/models"
	"github.com/dbzyuzin/swagger/internal/repostories/memory"
)

func CreateNewUser(ctx context.Context, user models.User) error {
	memory.AddUser(user)

	return nil
}

func UserExists(ctx context.Context, name string) (bool, error) {
	ok := memory.FindUser(name)
	return ok, nil
}
