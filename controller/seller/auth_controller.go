package seller_controller

import "github.com/labstack/echo/v4"

type SellerController interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
	View(c echo.Context) error
	Update(c echo.Context) error
}
