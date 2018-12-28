package auth

// Service provides user authentication operations.
type Service interface {
	AddUser(...User)
}

// Repository provides access to user repository.
type Repository interface {
	// AddUser saves a given user to the repository.
	AddUser(User) error
}

type service struct {
	bR Repository
}

// NewService creates an authentication service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// AddUser adds the given user(s) to the database
func (s *service) AddUser(b ...Beer) {
	// any validation can be done here
	for _, user := range b {
		_ = s.bR.AddUser(user) // error handling omitted for simplicity
	}
}
