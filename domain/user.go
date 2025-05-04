package domain

import "context"

type UserRepository interface {
	Create(ctx context.Context, user string) error
	GetUserByID(ctx context.Context, id string) (string, error)
}
