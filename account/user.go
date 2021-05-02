package account

import "context"

// User - to represent a user inside the business logic.
// json keys are to communicate via the HTTP transport layer.
type User struct {
	ID       string `json:"id,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Repository - interface to connect to the database, to put and get users to and from database.
type Repository interface {
	CreateUser(ctx context.Context, user User) error
	GetUser(ctx context.Context, id string) (string, error)
}