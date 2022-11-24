package v1

import "github.com/gogf/gf/v2/frame/g"

type UserRegisterReq struct {
	Email    string
	Nickname string
	Passport string
	Password string
	g.Meta   `path:"/register" tags:"User" method:"post" summary:"user register"`
}

type UserRegisterRes struct {
	g.Meta `mime:"application/json"`
}
