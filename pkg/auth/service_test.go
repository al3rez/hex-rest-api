package auth

import (
	"errors"
	"testing"
)

type testRepo struct {
	returnError error
}

func (t *testRepo) AddUser(u User) error {
	return t.returnError
}

func TestAddUser(t *testing.T) {
	t.Run("test returns error", func(t *testing.T) {
		s := NewService(&testRepo{returnError: errors.New("bad")})
		u := User{}
		err := s.AddUser(u)
		if err == nil {
			t.Fatal(err)
		}
	})

	t.Run("test adds user", func(t *testing.T) {
		s := NewService(&testRepo{})
		u := User{}
		err := s.AddUser(u)
		if err != nil {
			t.Fatal(err)
		}
	})
}
