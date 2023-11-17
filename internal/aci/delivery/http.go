package delivery

import (
	"github.com/labstack/echo/v4"
)

type ACIDelivery struct {
	api *echo.Group
}

func (a *ACIDelivery) InitRoute(route string, e *echo.Echo) error {
	api := e.Group(route)
	a.api = api
	return nil
}

func (a *ACIDelivery) BuildApi(e *echo.Echo) error {

	return nil
}

func NewACIDelivery() *ACIDelivery {
	return &ACIDelivery{}
}
