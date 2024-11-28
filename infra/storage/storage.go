package storage

import (
	"github.com/billykore/go-service-tmpl/domain/greet"
	"github.com/billykore/go-service-tmpl/infra/storage/repo"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	repo.NewGreetRepo, wire.Bind(new(greet.Repository), new(*repo.GreetRepo)),
)
