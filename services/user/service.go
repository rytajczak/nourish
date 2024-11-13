package main

import (
	"context"
	"user/repository"
)

type Service interface {
	CreateUser(ctx context.Context, user *repository.User) error
}
