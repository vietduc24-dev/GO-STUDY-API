package routers

import (
	"go-ecommerce-backend-api/internal/routers/manager"
	"go-ecommerce-backend-api/internal/routers/user"
)

type RouterGroup struct {
	User    user.UserRouterGroup
	Manager manager.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
