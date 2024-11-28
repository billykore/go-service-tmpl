package domain

import (
	"github.com/billykore/go-service-tmpl/domain/greet"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	greet.NewService,
)
