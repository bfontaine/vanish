package vanish

import (
	"io/ioutil"
	"os"
	"strings"
)

// File creates a temporary file and passes its name to the function passed as
// an argument. The file is deleted when the function returns.
func File(fn func(string)) error {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		return err
	}

	return callThenRemove(f.Name(), fn)
}

// Dir creates a temporary directory and passes its name to the function passed
// as an argument. The directory is deleted when the function returns.
func Dir(fn func(string)) error {
	name, err := ioutil.TempDir("", "")
	if err != nil {
		return err
	}

	return callThenRemove(name, fn)
}

// Env ensures the environment stays the same when the function passed as an
// argument is executed.
func Env(fn func()) error {
	env := os.Environ()

	fn()

	os.Clearenv()

	for _, pair := range env {
		kv := strings.SplitN(pair, "=", 2)
		os.Setenv(kv[0], kv[1])
	}

	return nil
}

func callThenRemove(name string, fn func(string)) error {
	defer os.Remove(name)

	fn(name)

	return nil
}
