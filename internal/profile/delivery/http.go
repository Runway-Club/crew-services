package delivery

import (
	profilePkg "core-crew.runwayclub.dev/core/domain/profile"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ProfileHttpDelivery struct {
	profileInterop profilePkg.ProfileInterop
	api            *echo.Group
}

func (p *ProfileHttpDelivery) Create(ctx echo.Context) error {
	token := ctx.Request().Header.Get("Authorization")
	profile := &profilePkg.Profile{}
	err := ctx.Bind(profile)
	if err != nil {
		return err
	}
	err = p.profileInterop.Create(ctx.Request().Context(), token, profile)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	return ctx.NoContent(http.StatusCreated)
}

func (p *ProfileHttpDelivery) InitRoute(route string, e *echo.Echo) error {
	p.api = e.Group(route)
	return nil
}

func (p *ProfileHttpDelivery) BuildApi(e *echo.Echo) error {
	p.api.POST("", p.Create)
	return nil
}

func NewProfileHttpDelivery(profileInterop profilePkg.ProfileInterop) *ProfileHttpDelivery {
	return &ProfileHttpDelivery{
		profileInterop: profileInterop,
	}
}
