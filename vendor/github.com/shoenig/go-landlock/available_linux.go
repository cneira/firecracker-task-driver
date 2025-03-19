// Copyright (c) Seth Hoenig
// SPDX-License-Identifier: MPL-2.0

//go:build linux

package landlock

var (
	available bool
	version   int
)

func init() {
	v, err := Detect()
	if err == nil {
		available = true
		version = v
	}
}

// Detect returns the version of landlock available, or an error.
func Detect() (int, error) {
	return abi()
}

// Available returns true if landlock is available, false otherwise.
func Available() bool {
	return available
}
