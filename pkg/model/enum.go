package model

import (
	"github.com/rafaelsanzio/go-core/pkg/applog"
	"github.com/rafaelsanzio/go-core/pkg/errs"
)

type Status string

var Statuses = make(map[string]Status, 3)

func status(name string) Status {
	i := Status(name)
	Statuses[name] = i
	return i
}

func GetStatus(key string) (Status, bool) {
	if key == "" {
		return "", false // default
	}

	val, ok := Statuses[key]
	if ok {
		return val, ok
	} else {
		return "", false
	}
}

func (i *Status) UnmarshalText(data []byte) error {
	str := string(data)

	val, ok := Statuses[str]

	if ok {
		*i = val
		return nil
	}

	return errs.ErrInvalidStatus.Throwf(applog.Log, "status: %s", str)
}

var (
	StatusActive    = status("Active")
	StatusInactive  = status("Inactive")
	StatusSuspended = status("Suspended")
)
