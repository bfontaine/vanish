# Vanish

**Vanish** is a minimal Go library to use temporary files and directories.

[![GoDoc](https://godoc.org/github.com/bfontaine/vanish?status.svg)](https://godoc.org/github.com/bfontaine/vanish/vanish)
[![Build Status](https://travis-ci.org/bfontaine/vanish.svg?branch=master)](https://travis-ci.org/bfontaine/vanish)

## Install

    go get github.com/bfontaine/vanish/vanish

## Usage

Vanish works with callbacks:

```go
import "github.com/bfontaine/vanish/vanish"

vanish.File(func(name string) {
    // 'name' is a temporary file, use it here as you want, it’ll be deleted
    // at the end of the function
})

vanish.Dir(func(name string) {
    // here, 'name' is a directory
})

vanish.Env(func() {
    // we can modify the environment here, it’ll be restored at the end of the
    // function
})
```
