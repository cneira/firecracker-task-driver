// Copyright (c) Seth Hoenig
// SPDX-License-Identifier: MPL-2.0

//go:build !linux

package landlock

import (
	"errors"
)

var (
	// ErrNotSupported indicates Landlock is not supported on this platform
	ErrNotSupported = errors.New("landlock not supported on this platform")

	// ErrBug indicates an unexpected error occured
	ErrBug = errors.New("landlock experienced an unexpected bug")
)

type locker struct {
	// does nothing
}

func New(...*Path) Locker {
	return new(locker)
}

func (l *locker) Lock(s Safety) error {
	switch s {
	case OnlyAvailable:
		return nil
	case OnlySupported:
		return nil
	case Try:
		return nil
	case Mandatory:
		return ErrNotSupported
	}
	return ErrBug
}

func (l *locker) String() string {
	return "[]"
}
