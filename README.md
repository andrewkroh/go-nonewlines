# go-nonewlines

[![Build Status](http://img.shields.io/travis/andrewkroh/go-nonewlines.svg?style=flat-square)][travis]
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)][godocs]

[travis]: http://travis-ci.org/andrewkroh/go-nonewlines
[godocs]: http://godoc.org/github.com/andrewkroh/go-nonewlines

go-nonewlines provides a `nonewlines` command that formats your Go code to
remove newlines that occur at the beginning and end of a function. It's not
a thoroughly tested tools so use with caution.

- Trim newlines between opening brace and first statement (except for multiline
  function declarations).
- Trim newlines before the closing brace.

## Installation and Usage

Package documentation can be found on [GoDoc][godocs].

Installation can be done with a normal `go get`.

```
$ go install github.com/andrewkroh/go-nonewlines/cmd/nonewlines
```

The arguments are similar to that of goimports. To view a diff
without actually updating any files use:

```
nonewlines -d .
```

To update files use

```
nonewlines -l -w .
```

## Example

```sh
$ nonewlines -d -l $(find . -name '*.go' | grep -v vendor) | head -20
./testify/assert/assertions.go
diff -u ./testify/assert/assertions.go.orig ./testify/assert/assertions.go
--- ./testify/assert/assertions.go.orig 2017-10-16 00:19:59.000000000 +0200
+++ ./testify/assert/assertions.go  2017-10-16 00:19:59.000000000 +0200
@@ -34,13 +34,11 @@
 //
 // This function does no assertion of any kind.
 func ObjectsAreEqual(expected, actual interface{}) bool {
-
    if expected == nil || actual == nil {
        return expected == actual
    }
 
    return reflect.DeepEqual(expected, actual)
-
 }
```
