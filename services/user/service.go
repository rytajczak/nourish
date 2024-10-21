package main

import "context"

type Service interface {
	Register(context.Context)
	Login(context.Context)
	CreateProfile(context.Context)
	GetProfile(context.Context)
	UpdateProfile(context.Context)
	DeleteProfile(context.Context)
}

type UserService struct{}

func (s *UserService) Register(context.Context) {
	panic("unimplemented")
}

func (s *UserService) Login(context.Context) {
	panic("unimplemented")
}

func (s *UserService) CreateProfile(context.Context) {
	panic("unimplemented")
}

func (s *UserService) GetProfile(context.Context) {
	panic("unimplemented")
}

func (s *UserService) UpdateProfile(context.Context) {
	panic("unimplemented")
}

func (s *UserService) DeleteProfile(context.Context) {
	panic("unimplemented")
}

func NewUserService() Service {
	return &UserService{}
}
