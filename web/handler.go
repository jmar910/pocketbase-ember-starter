package web

import (
	"embed"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

//go:embed ui/dist-prod/*
var distFS embed.FS

type Handler struct {
	app core.App
}

func NewHandler(app core.App) *Handler {
	return &Handler{app: app}
}

func (h *Handler) AddRoute(router *echo.Echo) {
	DistFS := echo.MustSubFS(distFS, "ui/dist-prod")
	router.GET(
		"/*",
		apis.StaticDirectoryHandler(DistFS, true),
		middleware.Gzip(),
	)
}
