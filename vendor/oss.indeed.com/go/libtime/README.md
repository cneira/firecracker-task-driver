libtime
=======

[![Go Report Card](https://goreportcard.com/badge/oss.indeed.com/go/libtime)](https://goreportcard.com/report/oss.indeed.com/go/libtime)
[![Build Status](https://travis-ci.com/indeedeng/libtime.svg?branch=master)](https://travis-ci.com/indeedeng/libtime)
[![GoDoc](https://godoc.org/oss.indeed.com/go/libtime?status.svg)](https://godoc.org/oss.indeed.com/go/libtime)
[![NetflixOSS Lifecycle](https://img.shields.io/osslifecycle/indeedeng/libtime.svg)](OSSMETADATA)
[![GitHub](https://img.shields.io/github/license/indeedeng/libtime.svg)](LICENSE)

# Project Overview

Package `libtime` is a library which provides common time operations.

# Getting Started

The `oss.indeed.com/go/libtime` module can be installed by running:
```
$ go get oss.indeed.com/go/libtime
```

Example usage:
```
c := libtime.SystemClock()

func foo(c libtime.Clock) {
    now := c.Now()
}
```

# Asking Questions

For technical questions about `libtime`, just file an issue in the GitHub tracker.

For questions about Open Source in Indeed Engineering, send us an email at
opensource@indeed.com

# Contributing

We welcome contributions! Feel free to help make `libtime` better.

### Process

- Open an issue and describe the desired feature / bug fix before making
changes. It's useful to get a second pair of eyes before investing development
effort.
- Make the change. If adding a new feature, remember to provide tests that
demonstrate the new feature works, including any error paths. If contributing
a bug fix, add tests that demonstrate the erroneous behavior is fixed.
- Open a pull request. Automated CI tests will run. If the tests fail, please
make changes to fix the behavior, and repeat until the tests pass.
- Once everything looks good, one of the indeedeng members will review the
PR and provide feedback.

# Maintainers

The `oss.indeed.com/go/libtime` module is maintained by Indeed Engineering.

While we are always busy helping people get jobs, we will try to respond to
GitHub issues, pull requests, and questions within a couple of business days.

# Code of Conduct

`oss.indeed.com/go/libtime` is governed by the [Contributer Covenant v1.4.1](CODE_OF_CONDUCT.md)

For more information please contact opensource@indeed.com.

# License

The `oss.indeed.com/go/libtime` module is open source under the [BSD-3-Clause](LICENSE) license.
