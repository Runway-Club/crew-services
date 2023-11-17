package delivery

import (
	"github.com/labstack/echo/v4"
)

type AuthDelivery struct {
	api *echo.Group
}

func NewAuthDelivery() *AuthDelivery {
	return &AuthDelivery{}
}

func (a *AuthDelivery) InitRoute(route string, e *echo.Echo) error {
	api := e.Group(route)
	a.api = api
	return nil
}

func (a *AuthDelivery) BuildApi(e *echo.Echo) error {

	return nil
}
