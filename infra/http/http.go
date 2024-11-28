package http

import (
	"github.com/billykore/go-service-tmpl/infra/http/handler"
	"github.com/billykore/go-service-tmpl/infra/http/server"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	handler.NewGreetHandler,
	server.NewRouter,
	server.New,
)
