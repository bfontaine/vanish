# Vanish

**Vanish** is a minimal Go library to use temporary files and directories.

## Install

    go get github.com/bfontaine/vanish/vanish

## Usage

Vanish works with callbacks:

```go
import "github.com/bfontaine/vanish/vanish"

vanish.File(function(name string) {
    // 'name' is a temporary file, use it here as you want, it’ll be deleted
    // at the end of the function
})

vanish.Dir(function(name string) {
    // here, 'name' is a directory
})

vanish.Env(function() {
    // we can modify the environment here, it’ll be restored at the end of the
    // function
})
```
