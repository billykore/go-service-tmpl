//go:build wireinject
// +build wireinject

package main

import (
	"github.com/billykore/go-service-tmpl/domain"
	"github.com/billykore/go-service-tmpl/infra/http"
	"github.com/billykore/go-service-tmpl/infra/storage"
	"github.com/billykore/go-service-tmpl/pkg"
	"github.com/billykore/go-service-tmpl/pkg/config"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

func initApp(cfg *config.Config) *app {
	wire.Build(
		domain.ProviderSet,
		storage.ProviderSet,
		http.ProviderSet,
		pkg.ProviderSet,
		echo.New,
		newApp,
	)
	return &app{}
}
