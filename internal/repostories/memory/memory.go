package memory

import (
	"context"
	"fmt"

	"github.com/dbzyuzin/swagger/internal/models"
	"github.com/dbzyuzin/swagger/internal/pkg/tracing"
)

type Repository struct {
	users []models.User
}

func (r *Repository) AddUser(name models.User) {
	r.users = append(r.users, name)
}

func (r *Repository) FindUser(ctx context.Context, name string) bool {
	fmt.Println("from repo", tracing.Get(ctx))
	for _, u := range r.users {
		if u.Name == name {
			return true
		}
	}

	return false
}
