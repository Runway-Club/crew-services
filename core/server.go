package core

import (
	"context"
	"github.com/labstack/echo/v4"
)

type Server struct {
	Echo *echo.Echo
}

func NewServer() *Server {
	e := echo.New()
	e.Group("/api/v1")
	return &Server{
		Echo: e,
	}
}

func Run(ctx context.Context, f func(ctx context.Context) error) error {
	return f(ctx)
}

func (s *Server) Start() error {
	return s.Echo.Start("0.0.0.0:8080")
}
