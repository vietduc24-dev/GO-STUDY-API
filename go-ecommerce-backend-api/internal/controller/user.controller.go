package controller

import (
	"go-ecommerce-backend-api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		UserService: service.NewUserService(),
	}
}

func (u *UserController) GetUserByID(c *gin.Context) {
	info := u.UserService.GetInfoUser()
	c.JSON(http.StatusOK, gin.H{
		"message": info,
	})
}
