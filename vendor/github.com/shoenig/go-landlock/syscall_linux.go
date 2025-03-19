// Copyright (c) Seth Hoenig
// SPDX-License-Identifier: MPL-2.0

//go:build linux

package landlock

import (
	"errors"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
	"kernel.org/pub/linux/libs/security/libcap/psx"
)

var (
	// ErrVersionUndetectable indicates a failure when checking landlock version
	ErrVersionUndetectable = errors.New("landlock version detection failure")

	// ErrNotSupported indicates landlock is not supported on this system
	ErrNotSupported = errors.New("landlock not supported")
)

type rule uint64

// FileSystem rules.
const (
	fsExecute     rule = unix.LANDLOCK_ACCESS_FS_EXECUTE
	fsWriteFile   rule = unix.LANDLOCK_ACCESS_FS_WRITE_FILE
	fsReadFile    rule = unix.LANDLOCK_ACCESS_FS_READ_FILE
	fsReadDir     rule = unix.LANDLOCK_ACCESS_FS_READ_DIR
	fsRemoveFile  rule = unix.LANDLOCK_ACCESS_FS_REMOVE_FILE
	fsRemoveDir   rule = unix.LANDLOCK_ACCESS_FS_REMOVE_DIR
	fsMakeChar    rule = unix.LANDLOCK_ACCESS_FS_MAKE_CHAR
	fsMakeDir     rule = unix.LANDLOCK_ACCESS_FS_MAKE_DIR
	fsMakeRegular rule = unix.LANDLOCK_ACCESS_FS_MAKE_REG
	fsMakeSocket  rule = unix.LANDLOCK_ACCESS_FS_MAKE_SOCK
	fsMakeFifo    rule = unix.LANDLOCK_ACCESS_FS_MAKE_FIFO
	fsMakeBlock   rule = unix.LANDLOCK_ACCESS_FS_MAKE_BLOCK
	fsMakeSymlink rule = unix.LANDLOCK_ACCESS_FS_MAKE_SYM
	fsRefer       rule = unix.LANDLOCK_ACCESS_FS_REFER
	fsTruncate    rule = unix.LANDLOCK_ACCESS_FS_TRUNCATE
)

func abi() (int, error) {
	r0, _, e1 := syscall.Syscall(
		unix.SYS_LANDLOCK_CREATE_RULESET,
		0,
		0,
		unix.LANDLOCK_CREATE_RULESET_VERSION,
	)
	v := int(r0)
	switch {
	case e1 != 0:
		return v, syscall.Errno(e1)
	case v < 0:
		return 0, ErrVersionUndetectable
	case v == 0:
		return 0, ErrNotSupported
	default:
		return v, nil
	}
}

func capabilities() rule {
	opts := fsExecute |
		fsWriteFile | fsReadFile | fsReadDir |
		fsRemoveFile | fsRemoveDir | fsMakeChar |
		fsMakeDir | fsMakeRegular | fsMakeSocket |
		fsMakeFifo | fsMakeBlock | fsMakeSymlink
	if version >= 2 {
		opts |= fsRefer
	}
	if version >= 3 {
		opts |= fsTruncate
	}
	return opts
}

type rulesetAttr struct {
	handleAccessFS uint64
}

func ruleset(ra *rulesetAttr) (int, error) {
	const size = 8
	r0, _, e1 := syscall.Syscall(
		unix.SYS_LANDLOCK_CREATE_RULESET,
		uintptr(unsafe.Pointer(ra)),
		uintptr(size),
		0,
	)
	return int(r0), errno(e1)
}

type beneathAttr struct {
	allowedAccess uint64
	parentFd      int
}

func add(fd int, ba *beneathAttr) error {
	_, _, e1 := syscall.Syscall6(
		unix.SYS_LANDLOCK_ADD_RULE,
		uintptr(fd),
		uintptr(unix.LANDLOCK_RULE_PATH_BENEATH),
		uintptr(unsafe.Pointer(ba)),
		0, 0, 0,
	)
	return errno(e1)
}

// https://git.kernel.org/pub/scm/libs/libcap/libcap.git/tree/psx/psx.go
//
// apply NO_NEW_PRIVS to all OS threads concurrently (with or without CGO)
func prctl() error {
	_, _, e1 := psx.Syscall6(
		syscall.SYS_PRCTL,
		unix.PR_SET_NO_NEW_PRIVS,
		1, 0, 0, 0, 0,
	)
	return errno(e1)
}

// https://git.kernel.org/pub/scm/libs/libcap/libcap.git/tree/psx/psx.go
//
// apply SYS_LANDLOCK_RESTRICT_SELF to all OS threads concurrently (with or without CGO)
func restrict(fd int) error {
	_, _, e1 := psx.Syscall3(
		unix.SYS_LANDLOCK_RESTRICT_SELF,
		uintptr(fd),
		0, 0,
	)
	return errno(e1)
}

func errno(e syscall.Errno) error {
	if e == 0 {
		return nil
	}
	return e
}
