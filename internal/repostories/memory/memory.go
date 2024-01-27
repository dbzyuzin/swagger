package memory

import "github.com/dbzyuzin/swagger/internal/models"

type Repository struct {
	users []models.User
}

func (r *Repository) AddUser(name models.User) {
	r.users = append(r.users, name)
}

func (r *Repository) FindUser(name string) bool {
	for _, u := range r.users {
		if u.Name == name {
			return true
		}
	}

	return false
}
