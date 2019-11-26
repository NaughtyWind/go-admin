package main

import (
	"go-admin/entity/auth"
)

func main() {
	u := new(auth.Register)
	u.LoginName = "lxy"
	u.PassWord = "12345"
	u.Phone = "123456"
	print(u.Phone)
}
