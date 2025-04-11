package usercontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/seaung/Dracula/pkg/types"
)

func (uc *UserController) Login(ctx *gin.Context) {
	var req types.UserLoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
	}

	ctx.JSON(200, gin.H{})
}
