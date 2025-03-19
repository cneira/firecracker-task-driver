// Copyright (c) Seth Hoenig
// SPDX-License-Identifier: MPL-2.0

// Package landlock provides a Go library for using the
// landlock feature of the modern Linux kernel.
//
// The landlock feature of the kernel is used to isolate
// a process from accessing the filesystem except for
// blessed paths and access modes.
package landlock

import (
	"fmt"
)

// Safety indicates the enforcement behavior on systems where landlock
// does not exist or operate as expected.
type Safety byte

const (
	// Mandatory mode will return an error on failure, including on
	// systems where landlock is not supported.
	Mandatory Safety = iota

	// OnlySupported will return an error on failure if running on a supported
	// operating system (Linux), or no error otherwise. Unlike OnlyAvailable,
	// this includes returning an error on systems where the Linux kernel was
	// built without landlock support.
	OnlySupported

	// OnlyAvailable will return an error on failure if running in an environment
	// where landlock is detected and available, or no error otherwise. Unlike
	// OnlySupported, OnlyAvailable does not cause an error on Linux systems built
	// without landlock support.
	OnlyAvailable

	// Try mode will continue with no error on failure.
	Try
)

// A Locker is an interface over the Kernel landlock LSM feature.
type Locker interface {
	fmt.Stringer
	Lock(s Safety) error
}
