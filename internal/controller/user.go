package controller

import (
	"context"
	"qa/internal/model"
	"qa/internal/service"

	v1 "qa/api/v1"
)

var (
	User = cUser{}
)

type cUser struct{}

func (c *cUser) SignUp(ctx context.Context, req *v1.UserRegisterReq) (res *v1.HelloRes, err error) {
	err = service.User().Create(ctx, model.UserCreateInput{
		Passport: req.Passport,
		Password: req.Password,
		Nickname: req.Nickname,
	})
	return
}
