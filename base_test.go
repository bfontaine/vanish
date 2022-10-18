package vanish

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func fileExists(path string) bool {
	// we don’t really check for errors here
	_, err := os.Stat(path)
	return err == nil
}

func TestFile(t *testing.T) {
	var filename string

	File(func(name string) {
		fi, err := os.Stat(name)
		assert.Nil(t, err)
		assert.True(t, fi.Mode().IsRegular())

		filename = name
	})

	assert.False(t, fileExists(filename))
}

func TestFileIn(t *testing.T) {
	Dir(func(dir string) {
		FileIn(dir, func(name string) {
			assert.True(t, strings.HasPrefix(name, dir))
			assert.True(t, fileExists(name))
		})
	})
}

func TestEmptyDir(t *testing.T) {
	var dirname string

	Dir(func(name string) {
		fi, err := os.Stat(name)
		assert.Nil(t, err)
		assert.True(t, fi.Mode().IsDir())

		dirname = name
	})

	assert.False(t, fileExists(dirname))
}

func TestDir(t *testing.T) {
	var dirname string

	assert.Nil(t, Dir(func(name string) {
		ioutil.TempFile(name, "")

		dirname = name
	}))

	assert.False(t, fileExists(dirname))
}

func TestDirIn(t *testing.T) {
	Dir(func(parent string) {
		DirIn(parent, func(dir string) {
			assert.True(t, fileExists(dir))
			assert.True(t, strings.HasPrefix(dir, parent))
		})
	})
}

func TestEnv(t *testing.T) {
	key := "VANISH_TEST_ENV_42XYZ"
	val := "foob&ar+/xqye$@"
	os.Setenv(key, val)

	Env(func() {
		assert.Equal(t, val, os.Getenv(key))
		os.Setenv(key, "yolofoo")
	})

	assert.Equal(t, val, os.Getenv(key))
}

func ExampleFile() {
	File(func(name string) {
		// 'name' is the name of a temporary file that’ll be deleted at the end
		// of the function. You can thus use it for your tests without worrying
		// about it
	})
}

func ExampleDir() {
	Dir(func(name string) {
		// Dir works like File, but it creates a directory instead. Use it as
		// you wish, it’ll be deleted at the end of this function.
	})
}

func ExampleEnv() {
	os.Setenv("VANISH", "xyz")

	Env(func() {
		// Env saves the environment before calling this function, ensuring
		// it’s restored as it was at its end.

		os.Setenv("VANISH", "foo")

		fmt.Printf("VANISH (1) = %s\n", os.Getenv("VANISH"))
	})

	fmt.Printf("VANISH (2) = %s\n", os.Getenv("VANISH"))

	// Output:
	// VANISH (1) = foo
	// VANISH (2) = xyz
}
