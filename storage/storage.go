package storage

import (
	"context"

	"github.com/luqxus/spaces/types"
)

type Storage interface {
	CreateUser(ctx context.Context, user *types.User) error
	CountEmail(ctx context.Context, email string) (int64, error)
	GetUserByEmail(ctx context.Context, email string) (*types.User, error)
}
