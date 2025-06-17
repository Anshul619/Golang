package repository

import (
    "context"
    "myservice/internal/model"
)

type UserRepository interface {
    FindByID(ctx context.Context, id string) (*model.User, error)
}
