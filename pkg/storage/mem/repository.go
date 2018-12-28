package mem

import (
	"time"

	"github.com/azbshiri/hex-rest-api/pkg/auth"
)

// Memory storage keeps data in memory
type Storage struct {
	users []User
}

// Add saves the given user to the repository
func (m *Storage) AddUser(u auth.User) error {
	newU := User{
		ID:      len(m.users) + 1,
		Name:    u.Name,
		Created: time.Now(),
	}
	m.users = append(m.users, newU)
	return nil
}
