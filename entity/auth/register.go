package auth

import (
	"go-admin/entity/basic"
)

type Register struct {
	basic.User
	Phone string `json:"phone"`
	Email string `json:"email"`
}
