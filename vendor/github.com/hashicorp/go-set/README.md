# go-set

[![Run CI Tests](https://github.com/hashicorp/go-set/actions/workflows/ci.yaml/badge.svg)](https://github.com/hashicorp/go-set/actions/workflows/ci.yaml)

Provides the `set` package that implements a generic mathematical [set](https://en.wikipedia.org/wiki/Set) for Go. The package only provides a basic implementation that is optimized for correctness and convenience. This package is not thread-safe.

# Documentation

The full documentation is available on GoDoc.

### Methods

Implements the following set operations

- Insert
- InsertAll
- InsertSet
- Remove
- RemoveAll
- RemoveSet
- Contains
- ContainsAll
- Size
- Union
- Difference
- Intersect

Provides helper methods

- Copy
- List
- String

# Install

```
go get github.com/hashicorp/go-set@latest
```

```
import "github.com/hashicorp/go-set"
```

# Example

Below are simple example of usages

```go
s := set.New[int](10)
s.Insert(1)
s.InsertAll([]int{2, 3, 4})
s.Size() # 3
```

```go
s := set.From[string]([]string{"one", "two", "three"})
s.Contains("three") # true
s.Remove("one") # true
```


```go
a := set.From[int]([]int{2, 4, 6, 8})
b := set.From[int]([]int{4, 5, 6})
a.Intersect(b) # {4, 6}
```
