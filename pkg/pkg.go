package pkg

import (
	"github.com/billykore/go-service-tmpl/pkg/logger"
	"github.com/billykore/go-service-tmpl/pkg/validation"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	logger.New,
	validation.New,
)
