package service

import (
	"context"
	"qa/internal/model"
)

type IUser interface {
	Create(ctx context.Context, in model.UserCreateInput) (err error)
}

var localUser IUser

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
