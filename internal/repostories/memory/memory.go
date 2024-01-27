package memory

import "github.com/dbzyuzin/swagger/internal/models"

var users []models.User

func AddUser(name models.User) {
	users = append(users, name)
}

func FindUser(name string) bool {
	for _, u := range users {
		if u.Name == name {
			return true
		}
	}

	return false
}
