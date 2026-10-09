package controller

import (
	"go-ecommerce-backend-api/internal/service"
	// "net/http"
    "go-ecommerce-backend-api/pkg/response"
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
	// info := u.UserService.GetInfoUser()
	// response.SuccessResponse(c, 2001, "Done")
	response.ErrorResponse(c, 2003, "Done")
}
