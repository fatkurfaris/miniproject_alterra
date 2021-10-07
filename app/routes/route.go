package routes

import (
	"kampus_merdeka/controllers/sellers"
	"kampus_merdeka/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JwtConfig        middleware.JWTConfig
	UserController   users.UserController
	SellerController sellers.SellerController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	e.POST("users/login", cl.UserController.Login)
	// e.GET("users", cl.UserController.Login, middleware.JWTWithConfig(cl.JwtConfig))
	e.POST("users/register", cl.UserController.Register)

	e.POST("sellers/login", cl.SellerController.Login)
	e.POST("sellers/register", cl.SellerController.Register)
}
