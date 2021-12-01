package config

import (
	"github.com/rafaelsanzio/go-core/pkg/config/key"
	"github.com/rafaelsanzio/go-core/pkg/errs"
)

type Service interface {
	Value(key.Key) (string, errs.AppError)
}
