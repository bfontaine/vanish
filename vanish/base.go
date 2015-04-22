// Vanish is a minimal Go library to use temporary files and directories.
package vanish

import (
	"io/ioutil"
	"os"
	"strings"
)

// File creates a temporary file and passes its name to the function passed as
// an argument. The file is deleted when the function returns.
func File(fn func(string)) error {
	return FileIn("", fn)
}

// FileIn is like File excepts that it accepts a parent directory as its first
// argument. If it's empty, the call is equivalent to File.
func FileIn(dir string, fn func(string)) error {
	f, err := ioutil.TempFile(dir, "")
	if err != nil {
		return err
	}

	f.Close()

	return callThenRemove(f.Name(), fn)
}

// Dir creates a temporary directory and passes its name to the function passed
// as an argument. The directory is deleted when the function returns.
func Dir(fn func(string)) error {
	return DirIn("", fn)
}

// DirIn is like Dir except that it accepts a parent directory as its first
// argument. If it’s empty, the call is equivalent to Dir.
func DirIn(dir string, fn func(string)) error {
	name, err := ioutil.TempDir(dir, "")
	if err != nil {
		return err
	}

	return callThenRemove(name, fn)
}

// Env ensures the environment stays the same when the function passed as an
// argument is executed.
func Env(fn func()) error {
	env := os.Environ()

	defer func() {
		os.Clearenv()

		for _, pair := range env {
			kv := strings.SplitN(pair, "=", 2)
			os.Setenv(kv[0], kv[1])
		}

	}()

	fn()

	return nil
}

func callThenRemove(name string, fn func(string)) (err error) {
	defer func() {
		err = os.RemoveAll(name)
	}()

	fn(name)

	return nil
}
