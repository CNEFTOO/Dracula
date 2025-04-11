package usercontroller

import (
	"github.com/seaung/Dracula/internal/App/biz"
	"github.com/seaung/Dracula/internal/App/store"
)

type UserController struct {
	b biz.Biz
}

func NewUserController(s store.Store) *UserController {
	return &UserController{b: biz.NewBiz(s)}
}
