package common

import "github.com/labstack/echo/v4"

type EchoDelivery interface {
	InitRoute(route string, e *echo.Echo) error
	BuildApi(e *echo.Echo) error
}

func NewHttpDelivery(root string, delivery EchoDelivery, e *echo.Echo) error {
	err := delivery.InitRoute(root, e)
	if err != nil {
		return err
	}
	err = delivery.BuildApi(e)
	if err != nil {
		return err
	}
	return nil
}
