# go-landlock

<img align="right" width="240" height="257" src="https://i.imgur.com/uOcXkpt.png">

[![Go Reference](https://pkg.go.dev/badge/github.com/shoenig/go-landlock.svg)](https://pkg.go.dev/github.com/shoenig/go-landlock)
[![CI Tests](https://github.com/shoenig/go-landlock/actions/workflows/ci.yaml/badge.svg)](https://github.com/shoenig/go-landlock/actions/workflows/ci.yaml)
[![License: MPL 2.0](https://img.shields.io/badge/License-MPL_2.0-orange.svg)](https://opensource.org/licenses/MPL-2.0)

The `go-landlock` module provides a Go library for interfacing with the Linux kernel
landlock feature. Landlock is a mechanism for minimizing filesystem access to a
Linux process. Using `go-landlock` does not require `root` or any escalated capabilities.

### Requirements

Requires **Linux 5.13+** with Landlock enabled. There is a no-op implementation provided
for non-Linux platforms for convenience, which provide no isolation.

Most recent Linux distributions should be supported.

Verified distros
- Ubuntu 22.04 LTS
- Ubuntu 20.04 LTS
- Fedora 36

The minimum Go version is `go1.19`.

### Install

Use `go get` to grab the latest version of `go-landlock`.

```shell
go get -u github.com/shoenig/go-landlock@latest
```

### Influence

This library is made possible after studying several sources, including but
not limited to

- [landlock.io](https://landlock.io/) official documentation
- [LWN](https://lwn.net/Articles/859908/)'s Landlock finally Sails
- [pledge.com](https://justine.lol/pledge/) by Justine Tunney
- [sandboxer.c](https://git.kernel.org/pub/scm/linux/kernel/git/stable/linux.git/tree/samples/landlock/sandboxer.c) kernel reference implementation

### API

Full documentation is on [pkg.go.dev](https://pkg.go.dev/github.com/shoenig/go-landlock).

The `go-landlock` package aims to provide a simple abstraction over the Kernel landlock
implementation details. Simply create a `Locker` with the `Path`'s to expose, and then
call `.Lock()` to isolate the process. The process will only be able to access the files
and directories, at the file modes specified. Attempts to access any other filesystem
paths will result in errors returned from the kernel system calls (like `open`).

Groups of commonly used paths are pre-defined for convenience.

- `Shared()` : for executing dynamically linked binaries
- `Stdio()` : for using standard I/O operations
- `TTY()` : for using terminal operations
- `Tmp()` : for accessing system tmp directory
- `VMInfo()` : for reading system information
- `DNS()` : for reading DNS related information
- `Certs()` : for reading system SSL/TLS certificate files

Custom paths can be specified using `File()` or `Dir()`. Each takes 2 arguments - the actual
filepath (absolute or relative), and a `mode` string. A mode string describes what level
of file mode permissions to allow. Must be a subset of `"rwxc"`.

- `r` : enable read permissions
- `w` : enable write permissions
- `x` : enable execute permissions
- `c` : enable create permissions

Once a `Locker` is configured, isolation starts on the call to `Lock()`. The level
of safety is configured by passing either `Mandatory` or `Try`.

- `Mandatory` : return an error is Landlock is unsupported or activation causes an error
- `Try` : continue without error regardless if landlock is supported or working
- `OnlySupported` : like `Mandatory`, but returns no error if the operating system does not support landlock

Once a process has been locked, it cannot be unlocked. Any descendent processes of the
locked process will also be locked, and cannot be unlocked. A child process can further
restrict itself via additional uses of landlock.

### Examples

#### complete example

This is a complete example of a small program which is able to read from
`/etc/os-release`, and is unable to access any other part of the filesystem

```go
package main

import (
  "fmt"
  "os"

  "github.com/shoenig/go-landlock"
)

func main() {
  l := landlock.New(
    landlock.File("/etc/os-release", "r"),
  )
  err := l.Lock(landlock.Mandatory)
  if err != nil {
    panic(err)
  }

  _, err = os.ReadFile("/etc/os-release")
  fmt.Println("reading /etc/os-release", err)

  _, err = os.ReadFile("/etc/passwd")
  fmt.Println("reading /etc/passwd", err)
}
```

```
➜ go run main.go
reading /etc/os-release <nil>
reading /etc/passwd open /etc/passwd: permission denied
```

#### shared objects (dynamic linking)

Programs that exec other processes may need to un-restrict a set of
shared object libraries. `go-landlock` provides the `Shared()` path
to simplify this configuration.

```go
l := landlock.New(
  landlock.Shared(), // common shared object files
  landlock.File("/usr/bin/echo", "rx"),
)

// e.g. execute echo in a sub-process
```

#### ssl/tls/dns (networking)

Programs that make use of the internet can use the `DNS()` and `Certs()`
helper paths to unlock necessary files for DNS resolution and reading
system SSL/TLS certificates.

```go
l := landlock.New(
  landlock.DNS(),
  landlock.Certs(),
)

// e.g.
// _, err = http.Get("https://example.com")
```

### License

Open source under the [MPL](LICENSE)
