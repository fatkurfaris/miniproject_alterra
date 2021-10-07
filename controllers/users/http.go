package users

import (
	"fmt"

	"kampus_merdeka/business/users"
	"kampus_merdeka/controllers"
	"kampus_merdeka/controllers/users/requests"
	"kampus_merdeka/controllers/users/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserUseCase users.Usecase
}

func NewUserController(userUseCase users.Usecase) *UserController {
	return &UserController{
		UserUseCase: userUseCase,
	}
}

func (userController UserController) Login(c echo.Context) error {

	userLogin := requests.UserLogin{}
	c.Bind(&userLogin)
	fmt.Println(userLogin)
	ctx := c.Request().Context()
	user, error := userController.UserUseCase.Login(ctx, userLogin.ToDomain())

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	// token, err := middlewares.CreateToken(users.Id )
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
	// 		"message": "fail login",
	// 		"error":   err.Error(),
	// 	})
	// }

	return controllers.NewSuccesResponse(c, responses.FromDomain(user))
}

func (userController UserController) Register(c echo.Context) error {
	userRegister := requests.UserRegister{}
	c.Bind(&userRegister)
	fmt.Println(userRegister)
	ctx := c.Request().Context()
	user, error := userController.UserUseCase.Register(ctx, userRegister.ToDomain())

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(user))
}
