package sellers

import (
	"fmt"
	"kampus_merdeka/business/sellers"
	"kampus_merdeka/controllers"
	"kampus_merdeka/controllers/sellers/requests"
	"kampus_merdeka/controllers/sellers/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SellerController struct {
	SellerUseCase sellers.Usecase
}

func NewSellerController(sellerUseCase sellers.Usecase) *SellerController {
	return &SellerController{
		SellerUseCase: sellerUseCase,
	}
}

func (sellerController SellerController) Login(c echo.Context) error {

	sellerLogin := requests.SellerLogin{}
	c.Bind(&sellerLogin)
	fmt.Println(sellerLogin)
	ctx := c.Request().Context()
	seller, error := sellerController.SellerUseCase.Login(ctx, sellerLogin.ToDomain())

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

	return controllers.NewSuccesResponse(c, responses.FromDomain(seller))
}

func (sellerController SellerController) Register(c echo.Context) error {
	sellerRegister := requests.SellerRegister{}
	c.Bind(&sellerRegister)
	fmt.Println(sellerRegister)
	ctx := c.Request().Context()
	seller, error := sellerController.SellerUseCase.Register(ctx, sellerRegister.ToDomain())

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(seller))
}
